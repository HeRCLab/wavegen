include ./opinionated.mk

PREFIX=/usr/local
BINDIR=$(PREFIX)/bin
MANDIR=$(PREFIX)/man


build: test build/bin/wavegen build/man/man1/wavegen.1 build/man/man1/wavegen-view.1 build/man/man1/wavegen-generate.1 build/man/man4/wavegen.4
.PHONY: build

test:
> go test -timeout 30s ./pkg/...
.PHONY: test

build/bin/wavegen:
> rm -f "$@"
> mkdir -p build/bin
> go build -ldflags="-s -w" -o "$@" ./cmd/wavegen/main.go
.PHONY: build/bin/wavegen

build/man/man1/wavegen.1: build/bin/wavegen
> mkdir -p build/man/man1
> help2man --include=include.txt --no-info --no-discard-stderr $< > "$@"
.PHONY: build/man/man1/wavegen

build/man/man4/wavegen.4: ./doc/format.md
> mkdir -p build/man/man4
> ronn < $< > $@

build/man/man1/wavegen-generate.1: build/bin/wavegen
> mkdir -p build/man/man1
> help2man --include=include.txt --no-info --no-discard-stderr "$< generate" > "$@"
.PHONY: build/man/man1/wavegen-generate.1

build/man/man1/wavegen-view.1: build/bin/wavegen
> mkdir -p build/man/man1
> help2man --include=include.txt --no-info --no-discard-stderr "$< view" > "$@"
.PHONY: build/man/man1/wavegen-view.1

install: build
> mkdir -p "$(BINDIR)"
> cp build/bin/* "$(BINDIR)"
> mkdir -p "$(MANDIR)/man1"
> mkdir -p "$(MANDIR)/man4"
> cp build/man/man1/* "$(MANDIR)/man1"
> cp build/man/man4/* "$(MANDIR)/man4"
.PHONY: install

fmt:
> go fmt ./pkg/...
.PHONY: fmt

lint:
> golint ./pkg/...
.PHONY: lint

coverage:
> go test -timeout 30s -coverprofile /dev/null ./pkg/...
.PHONY: coverage

viewcoverage:
> go test -timeout 30s -coverprofile cover.out ./pkg/...
> go tool cover -html=cover.out
.PHONY: viewcoverage

clean:
> rm -rf cover.out build release
.PHONY: clean
