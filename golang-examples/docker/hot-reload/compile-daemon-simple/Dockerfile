FROM golang:1.21.5-alpine3.18
RUN go install github.com/githubnemo/CompileDaemon@v1.4.0

WORKDIR /app

COPY ./ ./

ENTRYPOINT CompileDaemon --build="go build -o build/test_api_simple" -command="./build/test_api_simple" -build-dir=/app