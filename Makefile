.PHONY: lint build test

default: build

lint:
	go vet ./...

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mqtt-transformer ./cmd/main.go

