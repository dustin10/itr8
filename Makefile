all: test

format:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run ./...

test:
	go test -count=1 ./...

test-coverage:
	go test -count=1 ./... -coverprofile=coverage.out

test-coverage-view: test-coverage
	go tool cover -html=coverage.out

validate: sort-import format vet lint

.PHONY: test test-coverage test-coverage-view format vet lint validate
