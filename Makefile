go_packages := $(shell go list ./... | grep -v /vendor/)

.PHONY: default lint test

default:
	go run -race tik.go

lint:
	@golint -set_exit_status $(go_packages)

test: lint
	@go test -race $(go_packages)
