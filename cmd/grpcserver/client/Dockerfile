FROM golang:latest

ADD . /client
WORKDIR /client

ENTRYPOINT go run cmd/grpcserver/client/client.go

