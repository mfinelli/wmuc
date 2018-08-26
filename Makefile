SOURCES = $(wildcard **/*.go | grep -v ^vendor)

all: wmuc

clean:
	rm -rf vendor wmuc

fmt:
	find . -name 'vendor*' -prune -o -name '*.go' -exec go fmt {} \;

test: fmt vendor
	go test ./...

wmuc: $(SOURCES) vendor
	go build wmuc.go

vendor: Gopkg.toml Gopkg.lock
	dep ensure

.PHONY: all clean fmt test
