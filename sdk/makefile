SHELL := /bin/bash

# Variables
ROOT_DIR := $(shell pwd)
APP_NAME := golang-sdk
VERSION := 1.0
BUILD_DIR := ${ROOT_DIR}/bin
CONFIG_DIR := ${ROOT_DIR}/config
SCRIPTS_DIR := ${ROOT_DIR}/scripts

# Examples directory for APIs
EXAMPLES_DIR := /app/cmd/examples


# Phony targets
.PHONY: all build run test docker-up docker-down clean lint

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

# Docker targets for Monitoring API
monitoring-up:
	@echo "Starting Monitoring API in dev mode..."
	@chmod +x ${SCRIPTS_DIR}/entrypoint.sh
	APP_NAME=${APP_NAME} MAIN_DIR=${EXAMPLES_DIR}/monitoring/main.go docker compose -f ${CONFIG_DIR}/docker-compose.dev.yml --profile monitoring up

monitoring-build:
	@echo "Building Monitoring API in dev mode..."
	@chmod +x ${SCRIPTS_DIR}/entrypoint.sh
	APP_NAME=${APP_NAME} MAIN_DIR=${EXAMPLES_DIR}/monitoring/main.go docker compose -f ${CONFIG_DIR}/docker-compose.dev.yml --profile monitoring up --build

monitoring-down:
	@echo "Stopping Monitoring API in dev mode..."
	APP_NAME=${APP_NAME} MAIN_DIR=${EXAMPLES_DIR}/monitoring/main.go docker compose -f ${CONFIG_DIR}/docker-compose.dev.yml --profile monitoring down --remove-orphans

# Clean target
clean:
	@echo "Cleaning up..."
	@rm -f ${BUILD_DIR}/${APP_NAME}

# Lint target
lint:
	@echo "Linting the project..."
	@golangci-lint run --config .golangci.yml --verbose
