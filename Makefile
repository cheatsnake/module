.PHONY: test fmt vet all check

all: fmt vet test

check: fmt vet test

test:
	go test ./... -cover

fmt:
	go fmt ./...

vet:
	go vet ./...
