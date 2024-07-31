SHELL:=/bin/bash -O extglob
BINARY=num-man
VERSION=0.1
LDFLAGS=-ldflags "-X main.Version=${VERSION}"

#go tool commands
build:
	go build ${LDFLAGS} -o ${BINARY} cmd/main.go

run:
	@go run cmd/main.go

## tests
test:
	@go test ./...
	
## docker compose
up:
	docker-compose up --build -d
down:
	docker-compose down --remove-orphans