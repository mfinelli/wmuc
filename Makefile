SOURCES := $(wildcard *.go)
SOURCES += $(wildcard cmd/*.go)
SOURCES += $(wildcard lexer/*.go)
SOURCES += $(wildcard parser/*.go)
SOURCES += $(wildcard tokens/*.go)

all: wmuc

clean:
	rm -rf vendor wmuc legal/third_party.go

fmt:
	find . -name 'vendor*' -prune -o -name '*.go' -exec go fmt {} \;

test: fmt vendor
	go test ./...

wmuc: $(SOURCES) vendor legal/third_party.go
	go build wmuc.go

vendor: Gopkg.toml Gopkg.lock
	dep ensure

legal/third_party.go: scripts/license.go vendor
	go run scripts/license.go

.PHONY: all clean fmt test
