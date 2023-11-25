
.PHONY: all
all: test

.PHONY: test build
test:
	go test --failfast -v ./... -coverprofile=coverage.out -covermode=atomic -coverpkg=./...

.PHONY: build
build:
	go build -ldflags="-s -w" -o ./bin/ ./...

.PHONY: coverage-html
coverage-html:
	go tool cover -html=coverage.out

.PHONY: coverage-total
coverage-total:
	go tool cover -func=coverage.out | grep total | awk '{print $$3}'