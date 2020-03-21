cmake_minimum_required(VERSION 3.15)
project(SandBoxRunner C)

set(CMAKE_C_STANDARD 99)

add_executable(SandBoxRunner main.c rlimit.c rlimit.h seccmp.c seccmp.h child.c child.h)