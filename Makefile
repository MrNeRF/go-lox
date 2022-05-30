#https://go.dev/doc/code

TOKENS = pkg/tokens/*.go
UTILS = pkg/utils/*.go
SYNTAX = pkg/syntax-trees/*.go
PACKAGES = $(TOKENS) $(UTILS)

all: build install

build: $(PACKAGES)
	go build $(UTILS)
	go build $(TOKENS)
	go build $(SYNTAX)

install:
	go install
