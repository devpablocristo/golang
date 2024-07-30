#!/bin/zsh

protoc --go_out=. --go-grpc_out=. internal/proto/greet.proto