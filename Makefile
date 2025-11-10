# Makefile for Atlassian Data Center MCP

# Go compilation related variables
GOCMD ?= go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST  := $(GOCMD) test
GOMOD   := $(GOCMD) mod

# Version information
VERSION ?= $(shell git describe --tags --match "v*" 2>/dev/null | sed 's/^v//' || echo "dev")
COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
DATE ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)

# Build variables
BUILD_DIR     := dist
SERVER_BINARY := atlassian-dc-mcp-server
BINARY_PATH   := $(BUILD_DIR)/$(SERVER_BINARY)
SERVER_MAIN   := ./cmd/server

# Build flags
LDFLAGS := -ldflags="-s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)"
# Release build flags (static linking)
RELEASE_LDFLAGS := -ldflags="-s -w -extldflags '-static' -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)"

# Default target
.PHONY: all build build-linux build-windows build-macos release clean deps test run-server lint help
.DEFAULT_GOAL := help

all: build

# Clean build directory
clean:
	@echo "Cleaning build directory..."
	@rm -rf $(BUILD_DIR)

# Build server binary
build:
	@echo "Building server binary for current OS/ARCH..."
	@mkdir -p $(BUILD_DIR)
	@$(GOBUILD) -o $(BINARY_PATH) $(LDFLAGS) $(SERVER_MAIN)
	@echo "Server binary built at $(BINARY_PATH)"
	@echo "Binary size: $$(du -h $(BINARY_PATH) | cut -f1)"

# Cross-compilation targets
build-linux:
	@echo "Building for Linux (amd64)..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_PATH)-linux-amd64 $(LDFLAGS) $(SERVER_MAIN)
	@echo "Linux (amd64) server binary built at $(BINARY_PATH)-linux-amd64"
	@echo "Binary size: $$(du -h $(BINARY_PATH)-linux-amd64 | cut -f1)"

build-windows:
	@echo "Building for Windows (amd64)..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_PATH)-windows-amd64.exe $(LDFLAGS) $(SERVER_MAIN)
	@echo "Windows (amd64) server binary built at $(BINARY_PATH)-windows-amd64.exe"
	@echo "Binary size: $$(du -h $(BINARY_PATH)-windows-amd64.exe | cut -f1)"

build-macos:
	@echo "Building for macOS (amd64)..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_PATH)-darwin-amd64 $(LDFLAGS) $(SERVER_MAIN)
	@echo "macOS (amd64) server binary built at $(BINARY_PATH)-darwin-amd64"
	@echo "Binary size: $$(du -h $(BINARY_PATH)-darwin-amd64 | cut -f1)"

# Build statically linked release version
release:
	@echo "Building release binary (statically linked)..."
	@mkdir -p $(BUILD_DIR)
	@CGO_ENABLED=0 $(GOBUILD) -a -o $(BINARY_PATH) $(RELEASE_LDFLAGS) $(SERVER_MAIN)
	@echo "Release server binary built at $(BINARY_PATH)"
	@echo "Binary size: $$(du -h $(BINARY_PATH) | cut -f1)"

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@$(GOMOD) tidy

# Run tests
test:
	@echo "Running tests..."
	@$(GOTEST) -v ./...

# Run server in development mode
run-server:
	@echo "Running server in development mode..."
	@$(GOCMD) run $(SERVER_MAIN)

# Run linter
lint:
	@echo "Running linter..."
	@# Assuming golangci-lint is installed
	@golangci-lint run

# ==============================================================================
# Docker Targets
# ==============================================================================
DOCKER_IMAGE_NAME ?= ghcr.io/step-chen/atlassian-dc-mcp-go
DOCKER_IMAGE_TAG  ?= latest

.PHONY: docker-build docker-build-host-net docker-push docker-compose-up docker-compose-down docker-compose-nginx-up docker-compose-nginx-down docker-compose-traefik-up docker-compose-traefik-down

# Build the Docker image
docker-build:
	@echo "Building Docker image $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)..."
	@docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) .

# Build the Docker image with host network
docker-build-host-net:
	@echo "Building Docker image $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) with host network..."
	@docker build --network=host -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) .

# Push the Docker image to a registry
docker-push: docker-build
	@echo "Pushing Docker image $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)..."
	@docker push $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)

# Run services using docker-compose
docker-compose-up:
	@echo "Starting services with docker compose..."
	@IMAGE_NAME=$(DOCKER_IMAGE_NAME) docker compose up --build -d

# Stop services using docker-compose
docker-compose-down:
	@echo "Stopping services with docker compose..."
	@IMAGE_NAME=$(DOCKER_IMAGE_NAME) docker compose down

# Run services with Nginx using docker-compose
docker-compose-nginx-up:
	@echo "Starting services with Nginx and HTTPS..."
	@IMAGE_NAME=$(DOCKER_IMAGE_NAME) docker compose -f docker-compose.nginx.yml up -d

# Stop services with Nginx using docker-compose
docker-compose-nginx-down:
	@echo "Stopping services with Nginx..."
	@IMAGE_NAME=$(DOCKER_IMAGE_NAME) docker compose -f docker-compose.nginx.yml down

# Run services with Traefik using docker-compose
docker-compose-traefik-up:
	@echo "Starting services with Traefik and HTTPS..."
	@IMAGE_NAME=$(DOCKER_IMAGE_NAME) docker compose -f docker-compose.traefik.yml up -d

# Stop services with Traefik using docker-compose
docker-compose-traefik-down:
	@echo "Stopping services with Traefik..."
	@IMAGE_NAME=$(DOCKER_IMAGE_NAME) docker compose -f docker-compose.traefik.yml down

# Help information
help:
	@echo "Available commands:"
	@echo "  make all           Alias for 'make build'."
	@echo "  make build         Build the server binary for the current OS/ARCH."
	@echo "  make build-linux   Build the server binary for Linux (amd64)."
	@echo "  make build-windows Build the server binary for Windows (amd64)."
	@echo "  make build-macos   Build the server binary for macOS (amd64)."
	@echo "  make release       Build a statically linked release binary."
	@echo "  make run-server    Run the server for development."
	@echo "  make test          Run all tests with verbose output."
	@echo "  make deps          Tidy go.mod and go.sum."
	@echo "  make lint          Run the linter (requires golangci-lint)."
	@echo "  make clean         Remove the build directory."
	@echo ""
	@echo "Docker Commands:"
	@echo "  make docker-build              Build the Docker image."
	@echo "  make docker-build-host-net     Build the Docker image with host network."
	@echo "  make docker-push               Push the Docker image to a registry."
	@echo "  make docker-compose-up         Start services using docker compose."
	@echo "  make docker-compose-down       Stop services using docker compose."
	@echo "  make docker-compose-nginx-up   Start services with Nginx and HTTPS."
	@echo "  make docker-compose-nginx-down Stop services with Nginx."
	@echo "  make docker-compose-traefik-up   Start services with Traefik and HTTPS."
	@echo "  make docker-compose-traefik-down Stop services with Traefik."
	@echo ""
	@echo "  make help          Show this help message."