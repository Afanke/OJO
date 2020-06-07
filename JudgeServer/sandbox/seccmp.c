//
// Created by Afanke on 2020/2/13.
//

#include "seccmp.h"

int init_seccomp(){

    scmp_filter_ctx ctx;
    ctx = seccomp_init(SCMP_ACT_ALLOW);

    // as follows are the sys function that can be banned
#ifndef CPSBOX
    seccomp_rule_add(ctx, SCMP_ACT_KILL, SCMP_SYS(socket), 0);
     seccomp_rule_add(ctx, SCMP_ACT_KILL, SCMP_SYS(setuid), 0);
     seccomp_rule_add(ctx, SCMP_ACT_KILL, SCMP_SYS(creat), 0);
     seccomp_rule_add(ctx, SCMP_ACT_KILL, SCMP_SYS(fork), 0);
     seccomp_rule_add(ctx, SCMP_ACT_KILL, SCMP_SYS(clone), 0);
     seccomp_rule_add(ctx, SCMP_ACT_KILL, SCMP_SYS(vfork), 0);
     seccomp_rule_add(ctx, SCMP_ACT_KILL, SCMP_SYS(chmod), 0);
     seccomp_rule_add(ctx, SCMP_ACT_KILL, SCMP_SYS(open), 0);
     seccomp_rule_add(ctx, SCMP_ACT_KILL, SCMP_SYS(truncate), 0);
     seccomp_rule_add(ctx, SCMP_ACT_KILL, SCMP_SYS(writev), 0);
     seccomp_rule_add(ctx, SCMP_ACT_KILL, SCMP_SYS(s390_pci_mmio_write), 0);
     seccomp_rule_add(ctx, SCMP_ACT_KILL, SCMP_SYS(flock), 0);
     seccomp_rule_add(ctx, SCMP_ACT_KILL, SCMP_SYS(ftruncate), 0);
     seccomp_rule_add(ctx, SCMP_ACT_KILL, SCMP_SYS(umask), 0);
     seccomp_rule_add(ctx, SCMP_ACT_KILL, SCMP_SYS(fsync), 0);
     seccomp_rule_add(ctx, SCMP_ACT_KILL, SCMP_SYS(_llseek), 0);
     seccomp_rule_add(ctx, SCMP_ACT_KILL, SCMP_SYS(signal), 0);
#endif
    // as follows are the sys function that can not be banned
    //        seccomp_rule_add(ctx, SCMP_ACT_KILL, SCMP_SYS(lseek), 0);
    //        seccomp_rule_add(ctx, SCMP_ACT_KILL, SCMP_SYS(read), 0);
    //        seccomp_rule_add(ctx, SCMP_ACT_KILL, SCMP_SYS(write), 0);

    seccomp_load(ctx);
    return 0;
}
