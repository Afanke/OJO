all: CPS RTS

CPS:
	gcc main.c rlimit.c seccmp.c child.c -lseccomp -o CPSBOX -D CPSBOX
RTS:
	gcc main.c rlimit.c seccmp.c child.c -lseccomp -o RTSBOX