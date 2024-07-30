SHELL := /bin/bash

# Variables
ROOT_DIR := $(shell pwd)
APP_NAME := analytics
VERSION := 0.1
BUILD_DIR := ${ROOT_DIR}/bin

# Phony targets
.PHONY: all build run test docker-build docker-up docker-down clean lint

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

# Docker build target
docker-build:
	@echo "Building Docker images for production mode..."
	APP_NAME=${APP_NAME} docker-compose -f docker-compose.yml up --build

# Docker up target
docker-up:
	@echo "Starting Docker containers for production mode..."
	APP_NAME=${APP_NAME} docker-compose -f docker-compose.yml up

# Docker down target
docker-down:
	@echo "Stopping Docker containers for production mode..."
	APP_NAME=${APP_NAME} docker-compose -f docker-compose.yml down --remove-orphans

# Docker build target for debug mode
docker-dbuild:
	@echo "Building Docker images for debug mode..."
	APP_NAME=${APP_NAME} docker-compose -f docker-compose.debug.yml up --build

# Docker up target for debug mode
docker-dup:
	@echo "Starting Docker containers in debug mode..."
	APP_NAME=${APP_NAME} docker-compose -f docker-compose.debug.yml up

# Docker down target for debug mode
docker-ddown:
	@echo "Stopping Docker containers in debug mode..."
	APP_NAME=${APP_NAME} docker-compose -f docker-compose.debug.yml down --remove-orphans

# Clean target
clean:
	@echo "Cleaning up..."
	@rm -f ${BUILD_DIR}/${APP_NAME}

# Lint target
lint:
	@echo "Linting the project..."
	@golangci-lint run --config .golangci.yml --verbose
