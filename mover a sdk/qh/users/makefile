SHELL := /bin/bash -O extglob

# Read variables from the .env file
include .env
export

# Root directory of the project
ROOT_DIR := $(shell pwd)

# Main targets
.PHONY: all build run test docker-build docker-up docker-down clean lint

# Build and run the project
all: build run

# Build target
build:
	go build -gcflags "all=-N -l" -o ${ROOT_DIR}/app/main/bin/${BINARY_NAME} -ldflags "-X main.Version=${VERSION}" ${ROOT_DIR}/app/cmd/api/bin

# Run the project
run:
	@go run ${ROOT_DIR}/app/cmd/api/main.go

# Run tests
test:
	@go test ./...

# Build and bring up Docker Compose services
docker-build:
	docker-compose up --build

# Bring up Docker Compose services
docker-up:
	docker-compose up

# Stop and remove Docker Compose services
docker-down:
	docker-compose down --remove-orphans

# Clean binary files
clean:
	@rm -f ${ROOT_DIR}/app/main/bin/${BINARY_NAME}

# Linter for code validation
lint:
	@golangci-lint run
