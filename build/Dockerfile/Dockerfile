FROM ubuntu

RUN apt-get update -y
RUN apt-get -y upgrade
RUN apt-get install wget -y
RUN wget https://dl.google.com/go/go1.12.2.linux-amd64.tar.gz
RUN tar -xvf go1.12.2.linux-amd64.tar.gz
RUN mv go /usr/local
RUN export GOROOT=/usr/local/go
RUN export GOPATH=$HOME
RUN export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
RUN export GOBUILD=$HOME
RUN go get firebase.google.com/go
