FROM golang:1.20.1-alpine3.17 as build

# installing git
RUN apk update && apk upgrade && \
    apk add --no-cache git \
        dpkg \
        gcc \
        git \
        musl-dev
        
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN go get github.com/sirupsen/logrus
RUN go get github.com/buaazp/fasthttprouter
RUN go get github.com/valya la/fasthttp
# esta dependecia que se installa es para hacer el debug
RUN go get github.com/go-delve/delve/cmd/dlv

WORKDIR /go/src/work
COPY ./src /go/src/work/

RUN go build -o app
### Run the Delve debugger ###
COPY ./dlv.sh /
RUN chmod +x /dlv.sh 
ENTRYPOINT [ "/dlv.sh"]
# hasta aqui llega la image 'debug'
###########START NEW IMAGE###################

FROM alpine:3.17.2 as prod
COPY --from=debug /go/src/work/app /
CMD ./app
