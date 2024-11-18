SHELL := /bin/bash

GO_CMD=go
BUILD_DIR=build
BINARY_DIR=$(BUILD_DIR)/bin

deps: ## Install dependencies
	@echo "Installing dependencies..."
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	$(GO_CMD) get -u -t -d -v ./...
	$(GO_CMD) mod tidy
	$(GO_CMD) mod vendor

deps-cleancache: ## Clear cache in Go module
	@echo "Clearing cache in Go module..."
	$(GOCMD) clean -modcache

wire: ## Generate wire_gen.go
	@echo "Generating dependency injection..."
	wire gen chopipay/config/di

run: ## Run application
	@echo "Running application..."
	$(GO_CMD) run cmd/main.go

swag: ## Generate swagger documentation
	@echo "Generating swagger documentation..."
	swag init -g cmd/main.go