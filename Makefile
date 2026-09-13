BINARY_NAME=wx-api
BUILD_DIR=bin
MAIN_PATH=./cmd/wx-api/main.go
CONFIG_PATH=./config/config.yaml

.PHONY: all build run run-bin generate test test-race vet fmt lint clean tidy deps help

all: build

build: generate
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)

run: generate
	CONFIG_PATH=$(CONFIG_PATH) go run $(MAIN_PATH)

run-bin: build
	CONFIG_PATH=$(CONFIG_PATH) ./$(BUILD_DIR)/$(BINARY_NAME)

generate:
	go generate ./...

test:
	go test ./...

test-race:
	go test -race ./...

vet:
	go vet ./...

fmt:
	go fmt ./...

lint:
	golangci-lint run ./...

tidy:
	go mod tidy

deps:
	go mod download

clean:
	rm -rf $(BUILD_DIR)

help:
	@grep -E '^## ' $(MAKEFILE_LIST) | sed 's/## /  /'