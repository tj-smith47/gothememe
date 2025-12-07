.PHONY: all generate lint test test-short bench build clean prepare tools release-snapshot

all: generate lint test build

generate:
	go generate ./...
	go run ./cmd/themegen
	gofmt -w defaultthemes/

lint:
	golangci-lint run ./...

test:
	go test -v -race -coverprofile=coverage.out ./...

test-short:
	go test -v -short ./...

bench:
	go test -bench=. -benchmem ./...

build:
	go build -v ./...
	go build -v -o bin/themegen ./cmd/themegen

clean:
	rm -rf bin/ coverage.out
	rm -rf defaultthemes/*.go defaultthemes/*.svg
	rm -f default_registry.gen.go DEFAULT_THEMES.md

prepare:
	go mod tidy
	go mod verify

tools:
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.7.1
	go install golang.org/x/tools/cmd/goimports@latest
	go install golang.org/x/tools/cmd/stringer@latest

release-snapshot:
	goreleaser release --snapshot --clean

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

fmt:
	gofmt -w .
	goimports -w .
