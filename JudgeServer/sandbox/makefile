all: CPSBOX RTSBOX

CPSBOX:
	gcc main.c rlimit.c seccmp.c child.c -lseccomp -o CPSBOX -D CPSBOX
RTSBOX:
	gcc main.c rlimit.c seccmp.c child.c -lseccomp -o RTSBOX