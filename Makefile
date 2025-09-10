# Makefile for Atlassian Data Center MCP

# Go compilation related variables
GOCMD ?= go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST  := $(GOCMD) test
GOMOD   := $(GOCMD) mod

# Build variables
BUILD_DIR     := dist
SERVER_BINARY := atlassian-dc-mcp-server
BINARY_PATH   := $(BUILD_DIR)/$(SERVER_BINARY)
SERVER_MAIN   := ./cmd/server

# Build flags
LDFLAGS := -ldflags="-s -w"
# Release build flags (static linking)
RELEASE_LDFLAGS := -ldflags="-s -w -extldflags '-static'"

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
	@echo "Linux (amd64) binary built at $(BINARY_PATH)-linux-amd64"
	@echo "Binary size: $$(du -h $(BINARY_PATH)-linux-amd64 | cut -f1)"

build-windows:
	@echo "Building for Windows (amd64)..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_PATH)-windows-amd64.exe $(LDFLAGS) $(SERVER_MAIN)
	@echo "Windows (amd64) binary built at $(BINARY_PATH)-windows-amd64.exe"
	@echo "Binary size: $$(du -h $(BINARY_PATH)-windows-amd64.exe | cut -f1)"

build-macos:
	@echo "Building for macOS (amd64)..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_PATH)-darwin-amd64 $(LDFLAGS) $(SERVER_MAIN)
	@echo "macOS (amd64) binary built at $(BINARY_PATH)-darwin-amd64"
	@echo "Binary size: $$(du -h $(BINARY_PATH)-darwin-amd64 | cut -f1)"

# Build statically linked release version
release:
	@echo "Building release binary (statically linked)..."
	@mkdir -p $(BUILD_DIR)
	@CGO_ENABLED=0 $(GOBUILD) -a -o $(BINARY_PATH) $(RELEASE_LDFLAGS) $(SERVER_MAIN)
	@echo "Release binary built at $(BINARY_PATH)"
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
	@echo "  make help          Show this help message."