# Dockerfile

#the base image is fedora
FROM fedora:latest

# Build the image as root user
USER root

#installing powertop
RUN yum update && \
    yum install powertop && \

#starting powertop
CMD ["powertop"]
