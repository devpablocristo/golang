SHELL := /bin/bash

# Variables
ROOT_DIR := $(shell pwd)
APP_NAME := golang-sdk
VERSION := 1.0
BUILD_DIR := ${ROOT_DIR}/bin
CONFIG_DIR := ${ROOT_DIR}/config
SCRIPTS_DIR := ${ROOT_DIR}/scripts
EXAMPLES_DIR := /app/cmd/examples
DOCKER_COMPOSE_FILE := ${CONFIG_DIR}/docker-compose.dev.yml

# Phony targets
.PHONY: all build run test docker-up docker-build docker-down docker-logs clean lint greeter-server-up greeter-server-build greeter-server-down greeter-server-logs greeter-client-up greeter-client-build greeter-client-down greeter-client-logs monitoring-up monitoring-build monitoring-down go-micro-up go-micro-build go-micro-down

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
	@echo "Starting $(PROFILE) services in dev mode..."
	@chmod +x ${SCRIPTS_DIR}/entrypoint.sh
	APP_NAME=${APP_NAME} MAIN_DIR=${EXAMPLES_DIR}/$(PROFILE)/main.go docker compose -f ${DOCKER_COMPOSE_FILE} --profile $(PROFILE) up -d
	@$(MAKE) docker-logs PROFILE=$(PROFILE)

docker-build:
	@echo "Building $(PROFILE) services in dev mode..."
	@chmod +x ${SCRIPTS_DIR}/entrypoint.sh
	APP_NAME=${APP_NAME} MAIN_DIR=${EXAMPLES_DIR}/$(PROFILE)/main.go docker compose -f ${DOCKER_COMPOSE_FILE} --profile $(PROFILE) up --build -d
	@$(MAKE) docker-logs PROFILE=$(PROFILE)

docker-down:
	@echo "Stopping $(PROFILE) services in dev mode..."
	APP_NAME=${APP_NAME} MAIN_DIR=${EXAMPLES_DIR}/$(PROFILE)/main.go docker compose -f ${DOCKER_COMPOSE_FILE} --profile $(PROFILE) down --remove-orphans

docker-logs:
	@echo "Fetching logs for $(PROFILE) services..."
	@docker compose -f ${DOCKER_COMPOSE_FILE} logs -f $(SERVICE_NAME)

# Specific Docker targets using generic commands
greeter-server-up:
	@$(MAKE) docker-up PROFILE=greeter-server SERVICE_NAME=greeter-server-api

greeter-server-build:
	@$(MAKE) docker-build PROFILE=greeter-server SERVICE_NAME=greeter-server-api

greeter-server-down:
	@$(MAKE) docker-down PROFILE=greeter-server SERVICE_NAME=greeter-server-api

greeter-server-logs:
	@$(MAKE) docker-logs PROFILE=greeter-server SERVICE_NAME=greeter-server-api

greeter-client-up:
	@$(MAKE) docker-up PROFILE=greeter-client SERVICE_NAME=greeter-client-api

greeter-client-build:
	@$(MAKE) docker-build PROFILE=greeter-client SERVICE_NAME=greeter-client-api

greeter-client-down:
	@$(MAKE) docker-down PROFILE=greeter-client SERVICE_NAME=greeter-client-api

greeter-client-logs:
	@$(MAKE) docker-logs PROFILE=greeter-client SERVICE_NAME=greeter-client-api

monitoring-up:
	@$(MAKE) docker-up PROFILE=monitoring SERVICE_NAME=monitoring-api

monitoring-build:
	@$(MAKE) docker-build PROFILE=monitoring SERVICE_NAME=monitoring-api

monitoring-down:
	@$(MAKE) docker-down PROFILE=monitoring SERVICE_NAME=monitoring-api

go-micro-up:
	@$(MAKE) docker-up PROFILE=go-micro SERVICE_NAME=go-micro-api

go-micro-build:
	@$(MAKE) docker-build PROFILE=go-micro SERVICE_NAME=go-micro-api

go-micro-down:
	@$(MAKE) docker-down PROFILE=go-micro SERVICE_NAME=go-micro-api

go-micro-logs:
	@$(MAKE) docker-logs PROFILE=go-micro SERVICE_NAME=go-micro-api

# Clean target
clean:
	@echo "Cleaning up..."
	@rm -f ${BUILD_DIR}/${APP_NAME}

# Lint target
lint:
	@echo "Linting the project..."
	@golangci-lint run --config .golangci.yml --verbose
