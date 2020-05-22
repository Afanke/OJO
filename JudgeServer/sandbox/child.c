//
// Created by Afanke on 2020/2/13.
//


#include "child.h"

/*
 *  signal 14 : out of max_real_time
 *  signal 24 : out of max_cpu_time
 *  signal 11 : out of max_rss
 *  signal 31 : out of max_as
 */

void handle_child_sig(int sig){
    if (sig == SIGCHLD) {
        ftime(&tb_end);
        struct rusage res;
        int status=0;
        wait4(pid,&status,0,&res);
        if (WIFEXITED(status)) {
            fprintf(stderr,"^e%d", WEXITSTATUS(status));
        }
        if (WIFSIGNALED(status)) {
            fprintf(stderr,"^s%d", WTERMSIG(status));
        }
        fprintf(stderr,"m%ld", res.ru_maxrss);
        long sec = res.ru_stime.tv_sec + res.ru_utime.tv_sec;
        long usec = res.ru_stime.tv_usec + res.ru_utime.tv_usec;
        fprintf(stderr,"c%ld", sec*1000+usec/1000);
        flag = false;
    } else if (sig == SIGXCPU) {
        printf("sigcpu");
    } else {
        printf("%d",sig);
    }
}

void wait_for_child(){

    struct timeb start,end;
    ftime(&start);
    while (flag) {
        usleep(50000);
        ftime(&end);
        if (kill(pid, 0) == 0 ) {
//            printf("end-start= %ld\n", (end.time*1000+end.millitm) - (start.time*1000+start.millitm));
            if ((end.time*1000+end.millitm) - (start.time*1000+start.millitm)>=max_real_time) {
                kill(pid, 14);
//                    break;
            }
            continue;
        } else {
            break;
        }
    }
}