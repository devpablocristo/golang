FROM golang:latest
COPY simple.go .
RUN go build simple.go
EXPOSE 8080
CMD ["./simple"]