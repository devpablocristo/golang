SHELL := /bin/bash

# Variables
ROOT_DIR := $(shell pwd)
APP_NAME := golang-sdk
VERSION := 1.0
BUILD_DIR := ${ROOT_DIR}/bin
CONFIG_DIR := ${ROOT_DIR}/config
SCRIPTS_DIR := ${ROOT_DIR}/scripts

# Phony targets
.PHONY: all build run test docker-build docker-up docker-down docker-dbuild docker-dup docker-ddown clean lint

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

# NOTE: tuve cambiar de docker-compose a docker compose por un problema con la version de docker compose para ubuntu 24.04
# Docker build target
docker-build:
	@echo "Building Docker images for production mode..."
	APP_NAME=${APP_NAME} docker compose -f ${CONFIG_DIR}/docker-compose.yml up --build

# Docker up target
docker-up:
	@echo "Starting Docker containers for production mode..."
	APP_NAME=${APP_NAME} docker compose -f ${CONFIG_DIR}/docker-compose.yml up

# Docker down target
docker-down:
	@echo "Stopping Docker containers for production mode..."
	APP_NAME=${APP_NAME} docker compose -f ${CONFIG_DIR}/docker-compose.yml down --remove-orphans

# Docker build target for dev mode
docker-dbuild:
	@echo "Building Docker images for dev mode..."
	@chmod +x ${SCRIPTS_DIR}/entrypoint.sh
	APP_NAME=${APP_NAME} docker compose -f ${CONFIG_DIR}/docker-compose.dev.yml up --build

# Docker up target for dev mode
docker-dup:
	@echo "Starting Docker containers in dev mode..."
	@chmod +x ${SCRIPTS_DIR}/entrypoint.sh
	APP_NAME=${APP_NAME} docker compose -f ${CONFIG_DIR}/docker-compose.dev.yml up

# Docker down target for dev mode
docker-ddown:
	@echo "Stopping Docker containers in dev mode..."
	APP_NAME=${APP_NAME} docker compose -f ${CONFIG_DIR}/docker-compose.dev.yml down --remove-orphans

# Clean target
clean:
	@echo "Cleaning up..."
	@rm -f ${BUILD_DIR}/${APP_NAME}

# Lint target
lint:
	@echo "Linting the project..."
	@golangci-lint run --config .golangci.yml --verbose

# INFO: no funcionan estos comandos, revisar mas adelante
# Monitoring example
monitoring:
	@echo "Running monitoring example..."
	@chmod +x ${SCRIPTS_DIR}/entrypoint.sh
	APP_NAME=${APP_NAME} docker compose -f ${CONFIG_DIR}/docker-compose.dev.yml up --build
	APP_NAME=${APP_NAME} docker compose -f ${CONFIG_DIR}/docker-compose.dev.yml exec ${APP_NAME} go run ./cmd/examples/monitoring/main.go

# gRPC example
grpc:
	@echo "Running gRPC example..."
	docker compose -f ${CONFIG_DIR}/docker-compose.dev.yml run --rm --build golang-sdk go run ./cmd/examples/grpc/server/main.go

# Go-micro example
go-micro:
	@echo "Running go-micro example..."
	docker compose -f ${CONFIG_DIR}/docker-compose.dev.yml exec ${APP_NAME} go run ./cmd/examples/go-micro/main.go

# RabbitMQ example
rabbitmq:
	@echo "Running RabbitMQ example..."
	docker compose -f ${CONFIG_DIR}/docker-compose.dev.yml run --rm --build golang-sdk go run ./cmd/examples/rabbitmq/main.go

# INFO: fin comandos

# include .env.example
# export

# # HELP =================================================================================================================
# # This will output the help for each task
# # thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
# .PHONY: help

# help: ## Display this help screen
# 	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

# compose-up: ### Run docker compose
# 	docker compose up --build -d postgres rabbitmq && docker compose logs -f
# .PHONY: compose-up

# compose-up-integration-test: ### Run docker compose with integration test
# 	docker compose up --build --abort-on-container-exit --exit-code-from integration
# .PHONY: compose-up-integration-test

# compose-down: ### Down docker compose
# 	docker compose down --remove-orphans
# .PHONY: compose-down

# swag-v1: ### swag init
# 	swag init -g internal/controller/http/v1/router.go
# .PHONY: swag-v1

# run: swag-v1 ### swag run
# 	go mod tidy && go mod download && \
# 	DISABLE_SWAGGER_HTTP_HANDLER='' GIN_MODE=debug CGO_ENABLED=0 go run -tags migrate ./cmd/app
# .PHONY: run

# docker-rm-volume: ### remove docker volume
# 	docker volume rm my-hexagonal-go_pg-data
# .PHONY: docker-rm-volume

# linter-golangci: ### check by golangci linter
# 	golangci-lint run
# .PHONY: linter-golangci

# linter-hadolint: ### check by hadolint linter
# 	git ls-files --exclude='Dockerfile*' --ignored | xargs hadolint
# .PHONY: linter-hadolint

# linter-dotenv: ### check by dotenv linter
# 	dotenv-linter
# .PHONY: linter-dotenv

# test: ### run test
# 	go test -v -cover -race ./internal/...
# .PHONY: test

# integration-test: ### run integration-test
# 	go clean -testcache && go test -v ./integration-test/...
# .PHONY: integration-test

# mock: ### run mockery
# 	mockery --all -r --case snake
# .PHONY: mock

# migrate-create:  ### create new migration
# 	migrate create -ext sql -dir migrations 'migrate_name'
# .PHONY: migrate-create

# migrate-up: ### migration up
# 	migrate -path migrations -database '$(PG_URL)?sslmode=disable' up
# .PHONY: migrate-up
