FROM golang:latest


# ADD ../../ ./

RUN echo "hello"
RUN apt-get update
RUN apt-get install -y apt-utils
# RUN apt-get install -y mysql-client
RUN apt-get install -y vim
RUN apt-get install -y git
RUN go get bitbucket.org/liamstask/goose/cmd/goose
RUN git clone https://github.com/k-ueki/AGPluss.git

WORKDIR /go/src/AGPluss

# RUN go build

ENTRYPOINT ["echo","hello"]
