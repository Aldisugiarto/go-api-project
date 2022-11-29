FROM golang:alpine

WORKDIR /rest_api

ADD . .

RUN go mod download


RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT go build  && ./rest_api