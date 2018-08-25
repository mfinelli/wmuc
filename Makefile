SOURCES = $(wildcard **/*.go)

all: wmuc

fmt:
	find . -name '*.go' -exec go fmt {} \;

test: fmt
	go test ./...

wmuc: $(SOURCES)
	go build wmuc.go

.PHONY: fmt test
