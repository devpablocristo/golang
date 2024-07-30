FROM golang:1.20

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN make build
EXPOSE 8080
EXPOSE 8081
EXPOSE 8082
CMD ["./num-man"]