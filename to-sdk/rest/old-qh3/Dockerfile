# FROM golang:1.20

# RUN mkdir /app
# ADD . /app
# WORKDIR /app

# RUN make build
# EXPOSE 8080
# EXPOSE 8081
# EXPOSE 8082
# CMD ["./qh"]

# FROM golang:1.21.4-alpine3.18

# WORKDIR /go/src/app

# COPY . /go/src/app

# RUN go mod download

# RUN go build -o qh ./cmd/qh

# EXPOSE 8080

# CMD ["./qh"]


FROM golang:1.21.5-alpine3.18
RUN go install github.com/githubnemo/CompileDaemon@v1.4.0

WORKDIR /app

RUN apk update && \
    apk add libc-dev && \
    apk add gcc && \
    apk add make

COPY ./ /

COPY ./go.mod /go.mod
COPY ./go.sum /go.sum
RUN go mod download && go mod verify

COPY ./entrypoint.sh /entrypoint.sh
ADD https://raw.githubusercontent.com/eficode/wait-for/v2.2.4/wait-for /usr/local/bin/wait-for
RUN chmod +rx /usr/local/bin/wait-for /entrypoint.sh

ENTRYPOINT [ "sh", "/entrypoint.sh" ]