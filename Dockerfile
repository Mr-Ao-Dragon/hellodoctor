FROM debian:9
LABEL authors="mraodragon"
RUN rm /etc/apt/sources.list
RUN echo "deb https://archive.debian.org stretch main contrib non-free" > /etc/apt/sources.list
RUN echo "deb https://archive.debian.org stretch-updates main contrib non-free" > /etc/apt/sources.list
RUN echo "deb https://archive.debian.org stretch-backports main contrib non-free" > /etc/apt/sources.list.d/backports.list
RUN cat /etc/apt/sources.list \
    && apt-get update \
    && apt-get install -y go glibc-dev \
    && mkdir /src \
workdir /src