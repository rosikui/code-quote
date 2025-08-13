.PHONY: build clean test install help lint security release

# Binary name
BINARY_NAME=code-quote

# Build the application
build:
	go build -o $(BINARY_NAME) .

# Clean build artifacts
clean:
	go clean
	rm -f $(BINARY_NAME)
	rm -rf dist/

# Run tests
test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

# Run tests with coverage report
test-coverage: test
	go tool cover -html=coverage.txt -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Install dependencies
deps:
	go mod tidy
	go mod download

# Install the binary
install: build
	go install .

# Run the application
run: build
	./$(BINARY_NAME)

# Run linter
lint:
	golangci-lint run

# Run security scan
security:
	go install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck ./...

# Build for multiple platforms
build-all:
	mkdir -p dist
	GOOS=linux GOARCH=amd64 go build -o dist/code-quote-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build -o dist/code-quote-linux-arm64 .
	GOOS=darwin GOARCH=amd64 go build -o dist/code-quote-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build -o dist/code-quote-darwin-arm64 .
	GOOS=windows GOARCH=amd64 go build -o dist/code-quote-windows-amd64.exe .
	GOOS=windows GOARCH=arm64 go build -o dist/code-quote-windows-arm64.exe .

# Validate quotes
validate: build
	./$(BINARY_NAME) --lang en | head -1
	./$(BINARY_NAME) --daily | head -1
	./$(BINARY_NAME) --tag programming | head -1
	./$(BINARY_NAME) --markdown | head -1
	./$(BINARY_NAME) --no-color | head -1

# Show help
help:
	@echo "Available commands:"
	@echo "  build         - Build the application"
	@echo "  build-all     - Build for multiple platforms"
	@echo "  clean         - Clean build artifacts"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  deps          - Install dependencies"
	@echo "  install       - Install the binary"
	@echo "  run           - Build and run the application"
	@echo "  lint          - Run linter"
	@echo "  security      - Run security scan"
	@echo "  validate      - Validate built-in quotes"
	@echo "  help          - Show this help message"

# Default target
all: deps lint security test build validate