APP_NAME := xyz-finance-app

build:
	@echo "Building the applications..."
	go build -o bin/$(APP_NAME) cmd/main.go
	@echo "Build completed. Output: bin/$(APP_NAME)"

run: build
	@echo "Running the application..."
	./bin/$(APP_NAME)

run-dev:
	@echo "Running the application in development mode..."
	air

test:
	@echo "Running tests..."
	go test -v ./...
	@echo "Tests completed."

clean:
	@echo "Cleaning up build artifacts..."
	rm -f bin/$(APP_NAME)
	@echo "Build artifacts cleaned up."

tidy:
	@echo "Tidying up Go modules..."
	go mod tidy
	@echo "Go modules tidied up."

deps:
	@echo "Installing dependencies..."
	go mod download
	@echo "Dependencies installed."

help:
	@echo "Available commands:"
	@echo "  make build      Build the application"
	@echo "  make run        Build and run the application"
	@echo "  make clean      Clean up build artifacts"
	@echo "  make test       Run tests"
	@echo "  make tidy       Tidy up Go modules"
	@echo "  make deps       Install dependencies"
	@echo "  make help       Show this help message"
	@echo "  make run-dev    Run the application in development mode"