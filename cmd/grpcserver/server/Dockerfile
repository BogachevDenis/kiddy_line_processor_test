FROM golang:latest

ADD . /server
WORKDIR /server

ENTRYPOINT go run cmd/grpcserver/server/server.go

