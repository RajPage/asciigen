.PHONY: build test clean install run-example

BINARY_NAME=asciigen
BUILD_DIR=build

build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) .

test:
	go test ./... -v

test-coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

clean:
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html

install: build
	cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/

run-example: build
	@echo "Running example with sample image..."
	./$(BUILD_DIR)/$(BINARY_NAME) -input sample.jpg -width 80 -verbose

help:
	@echo "Available targets:"
	@echo "  build        - Build the binary"
	@echo "  test         - Run all tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  clean        - Clean build artifacts"
	@echo "  install      - Install binary to /usr/local/bin"
	@echo "  run-example  - Run with sample image"
	@echo "  help         - Show this help"
