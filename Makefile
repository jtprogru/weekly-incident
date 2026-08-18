.DEFAULT_GOAL := help
.PHONY: help build test cover lint fmt collect digest testdata clean

GO      ?= go
BIN     ?= bin
WEEK    ?=
DIGEST_ARGS := $(if $(WEEK),-week $(WEEK),)

help: ## Show this help
	@grep -hE '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2}'

build: ## Build both commands into ./bin
	$(GO) build -o $(BIN)/collect ./cmd/collect
	$(GO) build -o $(BIN)/digest ./cmd/digest

test: ## Run the test suite
	$(GO) test -race ./...

cover: ## Run tests with a coverage summary
	$(GO) test -coverprofile=cover.out ./...
	$(GO) tool cover -func=cover.out | tail -1

lint: ## Run go vet and golangci-lint
	$(GO) vet ./...
	golangci-lint run ./...

fmt: ## Format the tree
	gofmt -w .

collect: ## Fetch every source and merge into data/
	$(GO) run ./cmd/collect

digest: ## Render a week into weeks/ (WEEK=2026-W34 to pick one)
	$(GO) run ./cmd/digest $(DIGEST_ARGS)

testdata: ## Re-capture the golden feed responses used by the parser tests
	./scripts/fetch-testdata.sh

clean: ## Remove build output
	rm -rf $(BIN) cover.out
