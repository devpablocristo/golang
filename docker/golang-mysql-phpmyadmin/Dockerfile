FROM golang:1.14.4
RUN mkdir /app
COPY ./src /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]
#EXPOSE 15784 no hace nada