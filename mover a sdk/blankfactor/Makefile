# Cross parameters
SHELL:=/bin/bash -O extglob
BINARY=eve-srv
VERSION=0.1

LDFLAGS=-ldflags "-X main.Version=${VERSION}"

# Build step, generates the binary.
build: ## Build the binary
	@echo "Building $(BINARY)"
	@go build $(LDFLAGS) -o $(BINARY) cmd/main.go

# Run the web interface.
web: ## Run the web interface
	@echo "Running the web interface"
	@go run cmd/main.go

# Run the go formatter.
fmt: ## Run the go formatter
	@echo "Running the go formatter"
	@gofmt -w .

# Download the go lint.
lint-prepare: ## Download the go lint
	@echo "Installing golangci-lint"
	@curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

# Run the lint across all the project.
lint: ## Run the go lint
	@echo "Running the go lint"
	@./bin/golangci-lint run \
		--exclude-use-default=false \
		--enable=golint \
		--enable=gocyclo \
		--enable=goconst \
		--enable=unconvert \
		./...

# Run the test for all the directories.
test: ## Run the tests
	@echo "Running the tests"
	@go test -v ./...

###################
# Docker commands #
###################
up: ## Start the containers
	@echo "Starting the containers"
	@docker-compose up

bu: ## Start the containers
	@echo "Build or rebuild and start the containers"
	@docker-compose up --build

down: ## Stop the containers
	@echo "Stopping the containers"
	@docker-compose down --remove-orphans

clean: ## Clean up the data directory
	@echo "Cleaning up the data directory"
	@sudo rm -rf db/data

.PHONY: build web fmt lint-prepare lint test up down clean
