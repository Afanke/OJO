//
// Created by Afanke on 2020/2/10.
//
#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <string.h>
#include <wait.h>
#include <stdbool.h>
#include <bits/types/struct_rusage.h>
#include <sys/timeb.h>
#include "rlimit.h"
#include "seccmp.h"
#include "child.h"

pid_t pid;
volatile bool flag=true;
struct timeb tb_end;

int main(int argc, char *argv[])
{
    //接受参数
    if (argc < 6)
    {
        fprintf(stderr,"failed to load, lack of parameter");
        exit(0);
    }

    int max_cpu_time = (atoi(argv[1])/1000)+1;
    int max_real_time = atoi(argv[2]);
    int max_mem = atoi(argv[3]);
    char* stdin_path = argv[4];

    char** args = (char**)malloc(sizeof(char*)*(argc-5)+1);
    for (int i=0;i<(argc-5);i++){
        args[i]=argv[i+5];
    }
    args[(argc-5)]=(char*)0;

    //创建无名管道
    int pip1[2];
    int pip2[2];
    int pip3[2];
    if (pipe(pip1) == -1 || pipe(pip2) == -1 || pipe(pip3) == -1)
    {
        perror("failed to pipe");
        exit(0);
    }

    //创建子进程
    pid = fork();
    if (pid < 0)
    {
        perror("failed to fork");
        exit(0);
    }
    if (pid == 0)
    {
        if(strcmp(stdin_path,"0") != 0){
            freopen(stdin_path,"r",stdin);
        }
        dup2(pip1[1], 1); //dup stdout to pip1
        dup2(pip2[1], 2); //dup stderr to pip2
        struct timeb tb_start;
        ftime(&tb_start);
        write(pip3[1], &tb_start, sizeof(tb_start));
        init_seccomp();
        int res=0;
        if((res=init_rlimit(max_cpu_time, max_mem))!=0){
            fprintf(stderr,"failed to load rlimit, error:%d",res);
            exit(0);
        };
        execvp(args[0],args);
//        return 0;
    } else
    {
        char out_buff[1024 * 64] = {0};
        char err_buff[1024 * 64] = {0};
        struct timeb tb_start;
        signal(SIGCHLD, handle_child_sig);
        wait_for_child(max_real_time);
        close(pip1[1]);
        close(pip2[1]);
        close(pip3[1]);
        read(pip1[0], out_buff, sizeof(out_buff));
        read(pip2[0], err_buff, sizeof(err_buff));
        read(pip3[0], &tb_start, sizeof(tb_start));
        fprintf(stderr,"r%ld$", (tb_end.time*1000+tb_end.millitm) - (tb_start.time*1000+tb_start.millitm));
        printf("%s", out_buff);
        fprintf(stderr,"%s", err_buff);
    }
}
