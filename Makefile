SOURCES := $(wildcard *.go)
SOURCES += $(wildcard chuckfile/*.go)
SOURCES += $(wildcard cmd/*.go)
SOURCES += $(wildcard legal/*.go)
SOURCES += $(wildcard lexer/*.go)
SOURCES += $(wildcard parser/*.go)
SOURCES += $(wildcard tokens/*.go)
SOURCES += $(wildcard util/*.go)

PREFIX := /usr/local
DESTDIR :=

LDFLAGS := -ldflags '-s -w'

all: %.1 %.completion wmuc

clean:
	rm -rf vendor wmuc legal/third_party.go third-party.tar.gz* \
		wmuc-* wmuc.exe-* *.1 *.completion

fmt:
	find . -name 'vendor*' -prune -o -name '*.go' -exec go fmt {} \;

test: fmt vendor legal/third_party.go
	go test ./...

wmuc: $(SOURCES) vendor legal/third_party.go
	go build ${LDFLAGS} \
		-gcflags "all=-trimpath=${GOPATH}/src" \
		-asmflags "all=-trimpath=${GOPATH}/src" wmuc.go

vendor: Gopkg.toml Gopkg.lock
	dep ensure

legal/third_party.go: scripts/license/main.go vendor
	go run scripts/license/main.go

third-party.tar.gz: vendor
	tar zcvf third-party.tar.gz vendor

%.1: $(SOURCES) vendor legal/third_party.go
	go run scripts/doc/main.go

%.completion: $(SOURCES) vendor legal/third_party.go
	go run scripts/completions/main.go

install:
	install -Dm755 wmuc $(DESTDIR)$(PREFIX)/bin/wmuc
	install -Dm644 README.md $(DESTDIR)$(PREFIX)/share/doc/wmuc/README.md
	install -d $(DESTDIR)$(PREFIX)/share/man/man1
	install -m644 *.1 $(DESTDIR)$(PREFIX)/share/man/man1
	install -Dm644 bash.completion \
		$(DESTDIR)$(PREFIX)/share/bash-completion/completions/wmuc
	install -Dm644 zsh.completion \
		$(DESTDIR)$(PREFIX)/zsh/site-functions/_wmuc

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/wmuc
	rm -rf $(DESTDIR)$(PREFIX)/share/doc/wmuc
	rm -f $(DESTDIR)$(PREFIX)/share/man/man1/wmuc*.1
	rm -f $(DESTDIR)$(PREFIX)/share/bash-completion/completions/wmuc
	rm -f $(DESTDIR)$(PREFIX)/zsh/site-functions/_wmuc

release:
	./scripts/release.bash

.PHONY: all clean fmt install release test uninstall
