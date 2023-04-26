# Para hacer el rebuild mas eficiente
# es importante el orden de los comandos
# Poner primero las cosas que no cambian frecuentemente
# y luego las que cambiaran con cada build

# FROM golang:1.8

# RUN go get -d -v ./...
# RUN go install -v ./...

# COPY . /go/src/github.com/codeblooded/test1
# WORKDIR /go/src/github.com/codeblooded/test1

# RUN echo $PATH

# RUN go build -o test1 .
# CMD ["test1"]
# EXPOSE 3470


#When a layer change regarding previous build, 
#docker discards the following cached layers an runs them again, 
#sometimes wasting your time.

#Pay attention to the "Using cache" sentence that docker output in each layer 
#that is re-used from previous build.

#Another recommendation, for your dev work, use fresh to re-build your go app 
#automatically every time you change the code. Just installing it in the container and simply 
#using command: fresh in your docker-compose.yml

FROM golang:1.18.3-alpine3.16 as debug
# installing git
RUN apk update && apk upgrade && \
    apk add --no-cache git \
        dpkg \
        gcc \
        git 

#Todo esta esta parte es para el debber, probar si vale la pena mas adelante
#RUN apk update && apk upgrade && \
#    apk add --no-cache git \
#        dpkg \
#        gcc \
#        git 
#        git \
#        musl-dev 
        
#ENV GOPATH /go
#ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
#RUN go get github.com/sirupsen/logrus
#RUN go get github.com/buaazp/fasthttprouter
#RUN go get github.com/valya la/fasthttp
# esta dependecia que se installa es para hacer el debug
#RUN go get github.com/go-delve/delve/cmd/dlv

WORKDIR /app
COPY ./ /app

RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build src/server.go" --command=./server

EXPOSE 8080

###########START NEW IMAGE###################
# Los stages pueden agregarse asi y tambien en distintos dockerfiles
# Con scratch se puede hacer imagenes todavia menores
FROM alpine:3.16.0 as prod
COPY --from=debug /app /
CMD ./app
EXPOSE 8080
