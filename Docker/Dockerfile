FROM jenkins:2.60.3

USER root

RUN  apt-get update \
  && apt-get install -y wget \
  && rm -rf /var/lib/apt/lists/*

RUN apt-get update && apt-get install -y git

RUN wget https://dl.google.com/go/go1.11.2.linux-amd64.tar.gz

RUN tar -C /usr/local -xzf go1.11.2.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=/home/root/go
ENV GOROOT=/usr/local/go
ENV GOBIN=$GOPATH/bin
ENV PATH=$PATH:$GOPATH/bin
ENV CGO_ENABLED=0
