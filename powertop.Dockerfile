FROM fedora:latest

RUN yum update && \
    yum install powertop && \

RUN powertop
