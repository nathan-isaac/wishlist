# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	@templ generate
	@go build -o ./bin/main cmd/api/main.go

build_dev:
	@echo "Building..."
	@go build -o ./bin/main cmd/api/main.go

# Tailwind generation watch
tailwind:
	@echo "Generating Tailwind CSS..."
	@pnpm tailwindcss -i ./input.css -o ./cmd/web/assets/tailwind.css --watch

templ:
	@echo "Generating templates..."
	@templ generate -watch -proxy=http://localhost:8080

# Run the application
run:
	@go run cmd/api/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./tests -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f ./bin/main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: all build run test clean tailwind templ watch
