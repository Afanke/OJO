//
// Created by Afake on 2020/2/13.
//
#pragma once
//#ifndef SANDBOXRUNNER_CHILD_H
//#define SANDBOXRUNNER_CHILD_H
#include <wait.h>
#include <time.h>
#include <bits/types/struct_rusage.h>
#include <stdbool.h>
#include <stdio.h>
#include <unistd.h>
#include <sys/time.h>
#include <sys/timeb.h>

extern bool flag;
extern int pid;
extern struct timeb tb_end;
//#endif SANDBOXRUNNER_CHILD_H
void handle_child_sig(int sig);

void wait_for_child();
