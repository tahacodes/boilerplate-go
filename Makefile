# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod

# Target binary name
BINARY_PATH=./bin/application

# Main build target
all: clean build

# Prepare development environment
prepare:
	$(GOMOD) tidy
	$(GOMOD) vendor

# Build the binary
build:
	$(GOBUILD) -o $(BINARY_PATH) .

# Clean build artifacts
clean:
	$(GOCLEAN)
	rm -f $(BINARY_PATH)

# Run tests
test:
	$(GOTEST) ./...

# Run the binary
run:
	$(BINARY_PATH)
