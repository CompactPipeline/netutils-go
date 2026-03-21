BINARY=netutils-go
VERSION=0.1.0
LDFLAGS=-ldflags "-X main.version=$(VERSION)"

.PHONY: build test clean run lint

build:
	go build $(LDFLAGS) -o $(BINARY) .

test:
	go test -v -race ./...

coverage:
	go test -coverprofile=coverage.txt ./...
	go tool cover -html=coverage.txt

clean:
	rm -f $(BINARY) coverage.txt

lint:
	golangci-lint run

run: build
	./$(BINARY) https://httpbin.org/get https://httpbin.org/status/404