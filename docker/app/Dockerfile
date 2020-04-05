FROM golang:latest

RUN apt-get update
RUN apt-get install -y apt-utils
RUN apt-get install -y vim
RUN apt-get install -y git
RUN go get bitbucket.org/liamstask/goose/cmd/goose

WORKDIR /go/src/AGPluss

ENTRYPOINT ["echo","hello"]
