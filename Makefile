SOURCES := $(wildcard *.go)
SOURCES += $(wildcard cmd/*.go)
SOURCES += $(wildcard legal/*.go)
SOURCES += $(wildcard lexer/*.go)
SOURCES += $(wildcard parser/*.go)
SOURCES += $(wildcard tokens/*.go)
SOURCES += $(wildcard util/*.go)

all: wmuc

clean:
	rm -rf vendor wmuc legal/third_party.go third-party.tar.gz* \
		wmuc-* wmuc.exe-*

fmt:
	find . -name 'vendor*' -prune -o -name '*.go' -exec go fmt {} \;

test: fmt vendor legal/third_party.go
	go test ./...

wmuc: $(SOURCES) vendor legal/third_party.go
	go build wmuc.go

vendor: Gopkg.toml Gopkg.lock
	dep ensure

legal/third_party.go: scripts/license.go vendor
	go run scripts/license.go

third-party.tar.gz: vendor
	tar zcvf third-party.tar.gz vendor

release:
	./scripts/release.bash

.PHONY: all clean fmt release test
