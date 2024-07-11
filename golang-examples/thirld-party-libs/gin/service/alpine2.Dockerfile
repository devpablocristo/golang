ARG ALPINE_VERSION=3.13
ARG GO_VERSION=1.16
ARG BUILDPLATFORM=linux/amd64

FROM --platform=$BUILDPLATFORM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as builder

LABEL maintainer="Pablo Cristo, devpablocristo@gmail.com"

WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

# git install
RUN ["apk", "--update", "add", "git"]

COPY go.* ./
RUN ["go", "mod", "download"]
COPY . ./

RUN ["go", "get", "github.com/githubnemo/CompileDaemon@v1.2.1"]
RUN ["go", "get", "github.com/gin-gonic/gin@v1.7.1"]
ENTRYPOINT [`CompileDaemon -log-prefix=false -build="go build -o server main.go" -command="./server"`]