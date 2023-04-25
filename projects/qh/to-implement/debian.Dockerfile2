ARG VARIANT=1.16
ARG BUILDPLATFORM=linux/amd64
FROM --platform=$BUILDPLATFORM golang:${VARIANT} as builder

LABEL maintainer="Pablo Cristo, devpablocristo@gmail.com"

ENV DEBIAN_FRONTEND=noninteractive \
    GO111MODULE="on" \
    GOOS="linux" \
    CGO_ENABLED=0

RUN apt-get update
RUN apt-get -y install --no-install-recommends apt-utils 2>&1
RUN apt-get -y install git iproute2 procps lsb-release make openssh-client zsh

WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./

ARG AIR_VERSION=v1.27.3
ARG GIN_GONIC_VERSION=v1.7.1
ARG GOPLS_VERSION=v0.6.11
ARG GOPKGS_V2_VERSION=v2.1.2
ARG GO_OUTLINE_VERSION=v0.0.0-20200117021646-2a048b4510eb
ARG DLV_VERSION=v1.6.0
ARG STATICCHECK_VERSION=v0.1.4
RUN go get github.com/cosmtrek/air@${AIR_VERSION} \
    github.com/gin-gonic/gin@${GIN_GONIC_VERSION} \
    golang.org/x/tools/gopls@${GOPLS_VERSION} \
    github.com/uudashr/gopkgs/v2/cmd/gopkgs@${GOPKGS_V2_VERSION} \
    github.com/ramya-rao-a/go-outline@${GO_OUTLINE_VERSION} \
    github.com/go-delve/delve/cmd/dlv@${DLV_VERSION} \
    honnef.co/go/tools/cmd/staticcheck@${STATICCHECK_VERSION}

ENTRYPOINT ["air"]