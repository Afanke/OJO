FROM centos
RUN  mkdir judgeServer      \
&&   useradd judge          \
&& yum -y install gcc       \
&& cp /bin /judgeServer   -rf \
&& cp /lib /judgeServer   -rf \
&& cp /lib64 /judgeServer -rf \
&& cp /usr /judgeServer   -rf \
&& cp /sbin /judgeServer  -rf \
&& cp /etc /judgeServer   -rf 
COPY ./CPSBOX /judgeServer/
COPY ./RTSBOX /judgeServer/
COPY ./JudgeServer /judgeServer/
RUN chmod    0771 /judgeServer  \
&&  chmod -R 0775 /judgeServer/etc  \
&&  chmod -R 0775 /judgeServer/usr  \
&&  chmod    0775 /judgeServer/CPSBOX  \
&&  chmod    0775 /judgeServer/RTSBOX  \
&&  chmod    0775 /judgeServer/JudgeServer  
CMD  cd judgeServer \
&&   chroot . /bin/bash       

