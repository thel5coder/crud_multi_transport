GOCMD=go
BINARY_NAME=rli

all: clean lint test build

build:
	go build -o bin/$(BINARY_NAME) ./cmd/rli

clean:
	rm -rf bin

lint:
	golint ./...

test:
	go test ./...