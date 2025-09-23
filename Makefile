.PHONY: build test clean lint

# Build the application
build:
	go build -o bin/cyverApiCli cmd/main.go

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Run linter
lint:
	golangci-lint run

# Install dependencies
deps:
	go mod download

# Generate mocks
mocks:
	mockgen -source=internal/api/client.go -destination=internal/mocks/client_mock.go

# Run the application
run:
	go run cmd/main.go 