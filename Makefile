# Variables
BINARY_NAME=arena
BUILD_DIR=bin

.DEFAULT_GOAL := build

.PHONY: build clean test run tidy

## build: Compiles the application package into a binary
build: tidy
	@echo "🗜️ Building the binary..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/$(BINARY_NAME)/main.go

## run: Builds and executes the binary immediately
run: build
	@echo "🚀 Running $(BINARY_NAME)..."
	@./$(BUILD_DIR)/$(BINARY_NAME)


## tidy: Cleans up and downloads tracking dependencies 
tidy:
	@echo "📦 Tidying Go modules..."
	go mod tidy

## clean: Removes the generated build directory
clean:
	@echo "🧹 Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)