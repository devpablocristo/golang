# FROM golang:1.20

# RUN mkdir /app
# ADD . /app
# WORKDIR /app

# RUN make build
# EXPOSE 8080
# EXPOSE 8081
# EXPOSE 8082
# CMD ["./qh"]

FROM golang:1.21.4-alpine3.18

WORKDIR /go/src/app

COPY . /go/src/app

RUN go mod download

RUN go build -o crawler ./cmd/crawler

EXPOSE 8080

CMD ["./crawler"]
