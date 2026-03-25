BINARY_NAME=mac-monitor
GO_FILES=$(shell find . -name "*.go" -not -path "./vendor/*")

.PHONY: all build test clean run

all: build

build:
	CGO_ENABLED=1 go build -o bin/$(BINARY_NAME) main.go

test:
	go test -v -race ./...

clean:
	rm -rf bin/
	go clean

run: build
	./bin/$(BINARY_NAME)
