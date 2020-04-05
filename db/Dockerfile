FROM golang:latest

ENV DOCKERIZE_VERSION v0.6.1

RUN apt update
RUN apt-get install -y wget \
 && wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
 && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
 && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz
RUN go get -v github.com/ziutek/mymysql/...
RUN go get bitbucket.org/liamstask/goose/cmd/goose
