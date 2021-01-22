.PHONY: build test lint format

.DEFAULT:
build: test
	@goreleaser release --snapshot --rm-dist --skip-publish

test:
	@go test --cover -covermode=count -coverprofile=coverage.out ./...

lint:
	@golangci-lint run ./... -v

.ONESHELL:
format:
	@gofmt -l **/*.go
	@gofmt -l *.go
	@go fmt ./...
