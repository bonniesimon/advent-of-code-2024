# Define variables
APP_NAME := build/main
GO_FILES := main.go
INPUT_FILE := examples/sample-input.txt

# Default target (run)
run: build
	./$(APP_NAME) $(INPUT_FILE)

# Build the binary
build:
	go build -o $(APP_NAME) $(GO_FILES)

# Clean up generated files
clean:
	rm -f $(APP_NAME)

# Help message
help:
	@echo "Makefile for Go project"
	@echo "Usage:"
	@echo "  make         - Build and run the project"
	@echo "  make build   - Build the project"
	@echo "  make clean   - Clean the build artifacts"
	@echo "  make help    - Show this help message"

.PHONY: run build clean help
