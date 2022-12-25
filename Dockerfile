# syntax=docker/dockerfile:1

FROM golang:alpine

WORKDIR /rest_api

COPY . .
RUN go mod tidy

ENTRYPOINT go build  && ./rest_api

EXPOSE 3030

CMD [ "/golang" ]





