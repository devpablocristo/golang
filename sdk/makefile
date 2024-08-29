SHELL := /bin/bash

# Variables
ROOT_DIR := $(shell pwd)
APP_NAME := golang-sdk
VERSION := 1.0
BUILD_DIR := ${ROOT_DIR}/bin
CONFIG_DIR := ${ROOT_DIR}/config
SCRIPTS_DIR := ${ROOT_DIR}/scripts
EXAMPLES_DIR := /app/cmd/examples

# Phony targets
.PHONY: all build run test docker-up docker-build docker-down clean lint greeter-service-up greeter-service-build greeter-service-down monitoring-up monitoring-build monitoring-down go-micro-up go-micro-build go-micro-down

# Default target
all: build run

# Build target
build:
	@echo "Building the project..."
	@mkdir -p ${BUILD_DIR}
	go build -gcflags "all=-N -l" -o ${BUILD_DIR}/${APP_NAME} -ldflags "-X main.Version=${VERSION}" ${ROOT_DIR}/cmd

# Run target
run:
	@echo "Running the project..."
	@go run ${ROOT_DIR}/cmd/main.go

# Test target
test:
	@echo "Running tests..."
	@go test ./...

# Generic Docker commands
docker-up:
	@echo "Starting $(PROFILE) API in dev mode..."
	@chmod +x ${SCRIPTS_DIR}/entrypoint.sh
	APP_NAME=${APP_NAME} MAIN_DIR=${EXAMPLES_DIR}/$(PROFILE)/main.go docker compose -f ${CONFIG_DIR}/docker-compose.dev.yml --profile $(PROFILE) up -d

docker-build:
	@echo "Building $(PROFILE) API in dev mode..."
	@chmod +x ${SCRIPTS_DIR}/entrypoint.sh
	APP_NAME=${APP_NAME} MAIN_DIR=${EXAMPLES_DIR}/$(PROFILE)/main.go docker compose -f ${CONFIG_DIR}/docker-compose.dev.yml --profile $(PROFILE) up --build -d

docker-down:
	@echo "Stopping $(PROFILE) API in dev mode..."
	APP_NAME=${APP_NAME} MAIN_DIR=${EXAMPLES_DIR}/$(PROFILE)/main.go docker compose -f ${CONFIG_DIR}/docker-compose.dev.yml --profile $(PROFILE) down --remove-orphans

# Specific Docker targets using generic commands
greeter-server-up:
	@$(MAKE) docker-up PROFILE=greeter-server

greeter-server-build:
	@$(MAKE) docker-build PROFILE=greeter-server

greeter-server-down:
	@$(MAKE) docker-down PROFILE=greeter-server

greeter-client-up:
	@$(MAKE) docker-up PROFILE=greeter-client

greeter-client-build:
	@$(MAKE) docker-build PROFILE=greeter-client

greeter-client-down:
	@$(MAKE) docker-down PROFILE=greeter-client	
###
monitoring-up:
	@$(MAKE) docker-up PROFILE=monitoring

monitoring-build:
	@$(MAKE) docker-build PROFILE=monitoring

monitoring-down:
	@$(MAKE) docker-down PROFILE=monitoring

###
go-micro-up:
	@$(MAKE) docker-up PROFILE=go-micro

go-micro-build:
	@$(MAKE) docker-build PROFILE=go-micro

go-micro-down:
	@$(MAKE) docker-down PROFILE=go-micro
###

# Clean target
clean:
	@echo "Cleaning up..."
	@rm -f ${BUILD_DIR}/${APP_NAME}

# Lint target
lint:
	@echo "Linting the project..."
	@golangci-lint run --config .golangci.yml --verbose
