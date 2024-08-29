SHELL := /bin/bash

# Variables
ROOT_DIR := $(shell pwd)
APP_NAME := golang-sdk
VERSION := 1.0
BUILD_DIR := ${ROOT_DIR}/bin
CONFIG_DIR := ${ROOT_DIR}/config
SCRIPTS_DIR := ${ROOT_DIR}/scripts
EXAMPLES_DIR := /app/cmd/examples
DOCKER_COMPOSE := docker compose -f ${CONFIG_DIR}/docker-compose.dev.yml

# Phony targets
.PHONY: all build run test docker-up docker-build docker-down clean lint greeter-server-up greeter-server-build greeter-server-down monitoring-up monitoring-build monitoring-down go-micro-up go-micro-build go-micro-down

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
	APP_NAME=${APP_NAME} MAIN_DIR=${EXAMPLES_DIR}/$(PROFILE)/main.go ${DOCKER_COMPOSE} --profile $(PROFILE) up -d

docker-build:
	@echo "Building $(PROFILE) API in dev mode..."
	@chmod +x ${SCRIPTS_DIR}/entrypoint.sh
	APP_NAME=${APP_NAME} MAIN_DIR=${EXAMPLES_DIR}/$(PROFILE)/main.go ${DOCKER_COMPOSE} --profile $(PROFILE) up --build -d

docker-down:
	@echo "Stopping $(PROFILE) API in dev mode..."
	APP_NAME=${APP_NAME} MAIN_DIR=${EXAMPLES_DIR}/$(PROFILE)/main.go ${DOCKER_COMPOSE} --profile $(PROFILE) down --remove-orphans

# Specific Docker targets using generic commands
greeter-server-up: PROFILE = greeter-server
greeter-server-up: docker-up

greeter-server-build: PROFILE = greeter-server
greeter-server-build: docker-build

greeter-server-down: PROFILE = greeter-server
greeter-server-down: docker-down

greeter-client-up: PROFILE = greeter-client
greeter-client-up: docker-up

greeter-client-build: PROFILE = greeter-client
greeter-client-build: docker-build

greeter-client-down: PROFILE = greeter-client
greeter-client-down: docker-down

monitoring-up: PROFILE = monitoring
monitoring-up: docker-up

monitoring-build: PROFILE = monitoring
monitoring-build: docker-build

monitoring-down: PROFILE = monitoring
monitoring-down: docker-down

go-micro-up: PROFILE = go-micro
go-micro-up: docker-up

go-micro-build: PROFILE = go-micro
go-micro-build: docker-build

go-micro-down: PROFILE = go-micro
go-micro-down: docker-down

# Clean target
clean:
	@echo "Cleaning up..."
	@rm -f ${BUILD_DIR}/${APP_NAME}

# Lint target
lint:
	@echo "Linting the project..."
	@command -v golangci-lint >/dev/null 2>&1 || { echo >&2 "golangci-lint is not installed. Aborting."; exit 1; }
	@golangci-lint run --config .golangci.yml --verbose
