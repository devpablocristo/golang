#!/bin/zsh

# protoc --go_out=. --go_opt=paths=source_relative \
#     --go-grpc_out=. --go-grpc_opt=paths=source_relative \
#     internal/pb/chat.proto


protoc --go_out=. --go-grpc_out=. internal/pb/chat.proto