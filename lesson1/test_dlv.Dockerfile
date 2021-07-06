FROM centos
RUN yum install binutils -y \
 && yum install vim -y \
 && yum install golang -y \
 && yum install dlv -y \
 && yum install gdb -y
