FROM golang:1.21.5-alpine3.18

RUN apk update && \
    apk add --no-cache bash git libc-dev g++ inotify-tools && \
    go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /app

COPY . .

RUN go mod download && go mod verify

COPY ./entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

EXPOSE 8080 2345

# CMD ["dlv", "debug", "cmd/api/main.go", "--headless", "--listen=:2345", "--api-version=2", "--accept-multiclient", "--log"]
ENTRYPOINT ["/entrypoint.sh"]
