SOURCES = $(wildcard **/*.go)

all: wmuc

fetch:
	go get -v ./...

fmt:
	find . -name '*.go' -exec go fmt {} \;

test: fmt
	go test ./...

wmuc: $(SOURCES)
	go build wmuc.go

.PHONY: fetch fmt test
