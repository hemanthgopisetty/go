.DEFAULT_GOAL := build 
.PHONY: fmt vet build 

fmt: 
	@echo "Running go fmt..."
	@go fmt ./...

vet: fmt
	@echo "Running go vet..."
	@go vet ./...

build: vet
	@echo "Building Go application..."
	@go build