.DEFAULT_GOAL := help
.PHONY: help build test cover lint vet fmt fmt-check collect digest testdata clean

GO      ?= go
BIN     ?= bin
WEEK    ?=
SUMMARY ?=
WARNING ?=

DIGEST_ARGS  := $(if $(WEEK),-week $(WEEK),)
COLLECT_ARGS := $(if $(SUMMARY),-summary $(SUMMARY),) $(if $(WARNING),-warning $(WARNING),)

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

vet: ## Run go vet
	$(GO) vet ./...

fmt: ## Format the tree
	gofmt -w .

fmt-check: ## Fail when the tree is not gofmt-clean
	@unformatted="$$(gofmt -l .)"; \
	if [ -n "$$unformatted" ]; then \
		echo 'gofmt needed on:'; \
		echo "$$unformatted"; \
		exit 1; \
	fi

# The whole gate, cheapest check first, so a stray tab fails in seconds instead
# of after golangci-lint has walked the tree. CI runs this target and nothing
# else, so a green run here is a green run there.
lint: fmt-check vet ## Run gofmt -l, go vet and golangci-lint
	golangci-lint run ./...

collect: ## Fetch every source and merge into data/ (SUMMARY=, WARNING= capture the run notes)
	$(GO) run ./cmd/collect $(COLLECT_ARGS)

# Silent on purpose: cmd/digest prints the rendered week to stdout and the
# workflow reads the week from there, so an echoed recipe line would become
# part of it.
digest: ## Render a week into weeks/ (WEEK=2026-W34 to pick one)
	@$(GO) run ./cmd/digest $(DIGEST_ARGS)

testdata: ## Re-capture the golden feed responses used by the parser tests
	./scripts/fetch-testdata.sh

clean: ## Remove build output
	rm -rf $(BIN) cover.out
