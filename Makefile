GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=bin/main
LINTER=golangci-lint

.PHONY: all
all: test build

.PHONY: build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v

.PHONY: test
test:
	$(GOTEST) ./... -v

.PHONY: install
install:
	$(GOGET) ./...

.PHONY: test
	$(GOTEST) ./... -v

.PHONY: run
run:
	$(GOCMD) run main.go

.PHONY: lint
lint:
	$(LINTER) run
