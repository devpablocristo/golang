SHELL:=/bin/bash -O extglob
BINARY=crawler
VERSION=0.1
LDFLAGS=-ldflags "-X main.Version=${VERSION}"

# Directorio raíz del proyecto
ROOT_DIR := $(shell pwd)

# Ruta del archivo .env en el directorio raíz
ENV_FILE := $(ROOT_DIR)/.env

# Comandos
build:
	go build -o $(ROOT_DIR)/${BINARY}  -v ./cmd/crawler/

runbin:
	$(ROOT_DIR)/${BINARY} crawl $(url) 

runcmd:
	@go run ./cmd/crawler/crawler.go ./cmd/crawler/launcher.go crawl $(url) 

test:
	@go test ./...

up:
	sudo docker-compose up --build -d    

down:
	sudo docker-compose down --remove-orphans
	
logs:
	sudo docker-compose logs crawler  

