# Project name
APP_NAME = urlies

# Default target
all: build

# Build binary
build:
	go build -o $(APP_NAME) main.go

# Run (with default test input)
run:
	go run main.go -u https://example.com -m all -o output/

# Tidy up dependencies
deps:
	go mod tidy

# Clean build output
clean:
	rm -f $(APP_NAME)

# Rebuild from scratch
rebuild: clean deps build

.PHONY: all build run clean deps rebuild
