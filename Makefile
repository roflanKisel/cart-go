.PHONY: build
build:
	go build -o bin/main main.go

.PHONY: run
run:
	go run main.go

.PHONY: lint
lint:
	golangci-lint run
