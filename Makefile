#
# progress
#

BIN_NAME    := progress
BIN_VERSION := 1.0.0
IMAGE_NAME  := hoanhan101/progress

.PHONY: build
build:  ## Build the executable binary
	go build -o ${BIN_NAME} cmd/progress.go

.PHONY: clean
clean:  ## Remove temporary files and build artifacts
	go clean -v ./...
	rm -f ${BIN_NAME} coverage.out

.PHONY: cover
cover: test  ## Run unit tests and open the coverage report
	go tool cover -html=coverage.out

.PHONY: lint
lint:  ## Lint project source files
	golint -set_exit_status ./...

.PHONY: test
test:  ## Run unit tests
	go test -short -race -coverprofile=coverage.out -covermode=atomic ./...

.PHONY: version
version: ## Print the version
	@echo "${BIN_VERSION}"

.PHONY: help
help:  ## Print usage information
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

.DEFAULT_GOAL := help
