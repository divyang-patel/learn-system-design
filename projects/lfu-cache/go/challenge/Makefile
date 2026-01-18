# Standard Interface for all projects

.PHONY: run test clean install

# Default to help
help:
	@echo "Available commands:"
	@echo "  make install   - Install dependencies"
	@echo "  make run       - Run the demo (main)"
	@echo "  make test      - Run tests"
	@echo "  make clean     - Clean up build artifacts"

install:
	go mod tidy

run:
	@echo "This is a library implementation."
	@echo "Please use 'make test' to verify your solution."

test:
	go test -v ./src/...

clean:
	go clean
