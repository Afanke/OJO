//
// Created by Afanke on 2020/2/13.
//
#pragma once
#ifndef SANDBOXRUNNER_RLIMIT_H
#define SANDBOXRUNNER_RLIMIT_H
int init_rlimit(int max_cpu_time, int max_rss);
#endif //SANDBOXRUNNER_RLIMIT_H
#ifndef SYS_RESOURCE_H

#include <sys/resource.h>

#define SYS_RESOURCE_H
#endif //SYS_RESOURCE_H

#ifndef RLIMIT_ERROR
#define RLIMIT_ERROR

#define RLIMIT_CPU_ERROR -1
#define RLIMIT_RSS_ERROR -2
#define RLIMIT_LOCKS_ERROR -3
#define RLIMIT_MEMLOCK_ERROR -4
#define RLIMIT_MSGQUEUE_ERROR -5
#define RLIMIT_SIGPENDING_ERROR -6
#define RLIMIT_CORE_ERROR -7
#define RLIMIT_NOFILE_ERROR -8
#define RLIMIT_FSIZE_ERROR -9
#define RLIMIT_NPROC_ERROR -10

#endif //RLIMIT_ENNOR


#ifndef STDIO_H
#define STDIO_H

#include <stdio.h>

#endif //STDIO_H
