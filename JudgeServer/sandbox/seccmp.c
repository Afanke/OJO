//
// Created by Afanke on 2020/2/13.
//

#include "seccmp.h"

int init_seccomp()
{

    scmp_filter_ctx ctx;
    ctx = seccomp_init(SCMP_ACT_ALLOW);

    // --------------------------- process control -------------------------------
#ifndef CPSBOX
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(fork), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(vfork), 0);
#endif
//    seccomp_rule_add(ctx, SCMP_ACT_TRACE(1, SCMP_SYS(clone), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(capset), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(setsid), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(setpgid), 0);
    // --------------------------- process control -------------------------------

    // ------------------------------ file I/O -------------------------------------
#ifndef CPSBOX
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(creat), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(fsync), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(chmod), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(fchmod), 0);
#endif
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(rename), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(mknod), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(link), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(symlink), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(chown), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(fchown), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(lchown), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(chroot), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(rmdir), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(mount), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(umount), 0);
    // ------------------------------ file I/O -------------------------------------

    // ------------------------------ signal -------------------------------------
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(signal), 0);
    // ------------------------------ signal -------------------------------------

    // --------------------------- user & group ----------------------------------
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(setuid), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(setgid), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(setregid), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(setreuid), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(setresgid), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(setresuid), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(setfsgid), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(setfsuid), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(setgroups), 0);
    // --------------------------- user & group ----------------------------------

    // ----------------------------- sys control ---------------------------------
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(reboot), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(swapoff), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(adjtimex), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(alarm), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(setitimer), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(settimeofday), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(stime), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(vm86), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(delete_module), 0);
    // ----------------------------- sys control ---------------------------------

    // -------------------------------- socket -----------------------------------
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(bind), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(accept), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(send), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(sendto), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(sendmsg), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(recv), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(recvfrom), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(recvmsg), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(listen), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(select), 0);
    seccomp_rule_add(ctx, SCMP_ACT_TRAP, SCMP_SYS(sendfile), 0);
    // -------------------------------- socket -----------------------------------

    // --------------------- syscall can't be baned (Java & Python)-----------------------
    //    seccomp_rule_add(ctx, SCMP_ACT_LOG, SCMP_SYS(readdir), 0);
    //    seccomp_rule_add(ctx, SCMP_ACT_LOG, SCMP_SYS(socket), 0);
    //    seccomp_rule_add(ctx, SCMP_ACT_LOG, SCMP_SYS(socketpair), 0);
    //    seccomp_rule_add(ctx, SCMP_ACT_LOG, SCMP_SYS(connect), 0);
    //    seccomp_rule_add(ctx, SCMP_ACT_LOG, SCMP_SYS(execve), 0);
    //    seccomp_rule_add(ctx, SCMP_ACT_LOG, SCMP_SYS(write), 0);
    //    seccomp_rule_add(ctx, SCMP_ACT_LOG, SCMP_SYS(writev), 0);
    //    seccomp_rule_add(ctx, SCMP_ACT_LOG, SCMP_SYS(truncate), 0);
    //    seccomp_rule_add(ctx, SCMP_ACT_LOG, SCMP_SYS(ftruncate), 0);
    //    seccomp_rule_add(ctx, SCMP_ACT_LOG, SCMP_SYS(chdir), 0);
    //    seccomp_rule_add(ctx, SCMP_ACT_LOG, SCMP_SYS(fchdir), 0);
    //    seccomp_rule_add(ctx, SCMP_ACT_LOG, SCMP_SYS(getdents), 0);
    //    seccomp_rule_add(ctx, SCMP_ACT_LOG, SCMP_SYS(unlink), 0);
    //    seccomp_rule_add(ctx, SCMP_ACT_LOG, SCMP_SYS(readlink), 0);
    //    seccomp_rule_add(ctx, SCMP_ACT_LOG, SCMP_SYS(kill), 0);
    //    seccomp_rule_add(ctx, SCMP_ACT_LOG, SCMP_SYS(futex), 0);
    //    seccomp_rule_add(ctx, SCMP_ACT_LOG, SCMP_SYS(execve), 0);
    //    seccomp_rule_add(ctx, SCMP_ACT_LOG, SCMP_SYS(brk), 0);
    //    seccomp_rule_add(ctx, SCMP_ACT_LOG, SCMP_SYS(access), 0);
    //    seccomp_rule_add(ctx, SCMP_ACT_LOG, SCMP_SYS(readlink), 0);
    //    seccomp_rule_add(ctx, SCMP_ACT_LOG, SCMP_SYS(openat), 0);
    //    seccomp_rule_add(ctx, SCMP_ACT_LOG, SCMP_SYS(stat), 0);
    //    seccomp_rule_add(ctx, SCMP_ACT_LOG, SCMP_SYS(fstat), 0);
    //    seccomp_rule_add(ctx, SCMP_ACT_LOG, SCMP_SYS(mmap), 0);
    //    seccomp_rule_add(ctx, SCMP_ACT_LOG, SCMP_SYS(mprotect), 0);
    //    seccomp_rule_add(ctx, SCMP_ACT_LOG, SCMP_SYS(getpid), 0);
    //    seccomp_rule_add(ctx, SCMP_ACT_LOG, SCMP_SYS(munmap), 0);
    // --------------------- syscall can't be baned (Java & Python)-----------------------

    seccomp_load(ctx);
    return 0;
}
