//
// Created by Afanke on 2020/2/13.
//


#include "rlimit.h"

int init_rlimit(int max_cpu_time, int max_mem)
{
    struct rlimit R_CPU = {max_cpu_time, max_cpu_time + 2};
    struct rlimit R_RSS = {max_mem * 1024, max_mem * 1024 + 20000};
//    struct rlimit R_AS = {max_mem * 1024 * 5, max_mem * 1024 * 5};
# ifdef CPSBOX
    struct rlimit R_FSIZE = {20*1024*1024, 20*1024*1024};
    struct rlimit R_NOFILE = {30, 30};
# endif
# ifndef CPSBOX
    struct rlimit R_FSIZE = {0, 0};
    struct rlimit R_NOFILE = {11, 11};
#endif
    struct rlimit R_LOCKS = {0, 0};
    struct rlimit R_MEMLOCK = {0, 0};
    struct rlimit R_MSGQUEUE = {0, 0};
    struct rlimit R_SIGPENDING = {0, 0};
    struct rlimit R_CORE = {0, 0};
    struct rlimit R_NPROC = {400, 400};

    if (setrlimit(RLIMIT_CPU, &R_CPU))
    {
        perror("failed in RLIMIT_CPU");
        return RLIMIT_CPU_ERROR;
    }
//    if (setrlimit(RLIMIT_AS, &R_AS))
//    {
//        perror("failed in RLIMIT_AS");
//        return RLIMIT_RSS_ERROR;
//    }
    if (setrlimit(RLIMIT_RSS, &R_RSS))
    {
        perror("failed in RLIMIT_RSS");
        return RLIMIT_RSS_ERROR;
    }

    if (setrlimit(RLIMIT_LOCKS, &R_LOCKS))
    {
        perror("failed in RLIMIT_LOCKS");
        return RLIMIT_LOCKS_ERROR;
    }
    if (setrlimit(RLIMIT_MEMLOCK, &R_MEMLOCK))
    {
        perror("failed in RLIMIT_MEMLOCK");
        return RLIMIT_MEMLOCK_ERROR;
    }
    if (setrlimit(RLIMIT_MSGQUEUE, &R_MSGQUEUE))
    {
        perror("failed in RLIMIT_MSGQUEUE");
        return RLIMIT_MSGQUEUE_ERROR;
    }

    if (setrlimit(RLIMIT_SIGPENDING, &R_SIGPENDING))
    {
        perror("failed in RLIMIT_SIGPENDING");
        return RLIMIT_SIGPENDING_ERROR;
    }
    if (setrlimit(RLIMIT_CORE, &R_CORE))
    {
        perror("failed in RLIMIT_CORE");
        return RLIMIT_CORE_ERROR;
    }
    if (setrlimit(RLIMIT_FSIZE, &R_FSIZE))
    {
        perror("failed in RLIMIT_FSIZE");
        return RLIMIT_FSIZE_ERROR;
    }
    if (setrlimit(RLIMIT_NOFILE, &R_NOFILE))
    {
        perror("failed in RLIMIT_NOFILE");
        return RLIMIT_NOFILE_ERROR;
    }
    if (setrlimit(RLIMIT_NPROC, &R_NPROC))
    {
        perror("failed in RLIMIT_NPROC");
        return RLIMIT_NPROC_ERROR;
    }
    return 0;
}