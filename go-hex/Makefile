

# Project settings
APP_NAME := go-hex # change me

# Go parameters
GO_FILES := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: all build run test tidy fmt lint clean

# Default target
all: build

# Build the binary
build:
	go build -o bin/$(APP_NAME) ./cmd/server/main.go

# Run the app (useful for small projects)
run-dev:
	GO_ENV=dev go run ./cmd/server/main.go

run-dev-watch:
	GO_ENV=dev air -c .air.toml

run-prod:
	make build && GO_ENV=prod ./bin/$(APP_NAME)

# Run tests
test:
	go test ./...

# Ensure dependencies are tidy
tidy:
	go mod tidy

# Format code
fmt:
	golangci-lint fmt

# Lint code (requires golangci-lint or similar tool)
lint:
	golangci-lint run

# Clean build artifacts
clean:
	rm -rf bin/
