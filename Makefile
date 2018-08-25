all: test

fmt:
	find . -name '*.go' -exec go fmt {} \;

test: fmt
	go test ./...

.PHONY: fmt test
