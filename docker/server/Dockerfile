FROM golang:latest

ADD . .

WORKDIR /go/src/AGPluss/server
RUN go get github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm
RUN ls /go/src/github.com
RUN go get github.com/k-ueki/AGPlus/server


RUN go build
