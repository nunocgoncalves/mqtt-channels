.PHONY: fmt lint test run build all

default: test

fmt:
	go fmt ./...

lint:
	go vet ./...

test:
	go test ./...

run:
	go run cmd/main.go

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mqtt-channels ./cmd/main.go

all: fmt lint test build

