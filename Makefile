MAKEFLAGS := --no-print-directory

default: help

help:
	@echo "Please use 'make <target>' where <target> is one of"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z\._-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

r: run
run: fmt ## Run the program, alias: r
	go run cmd/server/main.go

swagger-gen: ## Generate swagger documentation from the code's docstrings
	swag init -g ./main.go --output swagger --parseDependency

test: ## Run unit tests, alias: t
	go test ./... --cover -timeout=60s -parallel=64

fmt: ## Format go code using gofumpt, go mod tidy and golangci-lint --fix
	@go mod tidy
	@gofumpt -l -w .
	@golangci-lint run --fix --timeout 120s

install-tools: ## Install extra tools for development
	go install github.com/swaggo/swag/cmd/swag@master
	go install mvdan.cc/gofumpt@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint: ## Lint the code locally
	golangci-lint run --timeout 120s

