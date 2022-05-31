#https://go.dev/doc/code

TOKENS = pkg/tokens/*.go
UTILS = pkg/utils/*.go
PARSER = pkg/parser/*.go
PACKAGES = $(TOKENS) $(UTILS) $(PARSER)

all: build install

build: $(PACKAGES)
	go build $(UTILS)
	go build $(TOKENS)
	go build $(PARSER)

install:
	go install
