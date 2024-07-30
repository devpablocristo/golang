#!/bin/zsh

protoc --go_out=. --go-grpc_out=. internal/greeter/proto/greet.proto
protoc --go_out=. --go-grpc_out=. internal/prompt/proto/teamcubot.proto