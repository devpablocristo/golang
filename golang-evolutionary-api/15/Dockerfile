FROM golang:1.22.3-alpine3.20

WORKDIR /app

COPY . .

RUN go build -o /app/bin/myapp ./cmd/rest/main.go

EXPOSE 8080

CMD ["/app/bin/myapp"]
