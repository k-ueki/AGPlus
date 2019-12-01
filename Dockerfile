FROM golang:latest

# RUN apt-get install -y mysql-client
RUN go get bitbucket.org/liamstask/goose/cmd/goose
RUN apt-get update
RUN apt-get install -y vim

WORKDIR /go/src/AGPluss

# RUN go build 
# CMD ["make","migrate-up"]

