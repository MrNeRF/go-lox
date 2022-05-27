#https://go.dev/doc/code

TOKENS = pkg/tokens/*.go
UTILS = pkg/utils/*.go
PACKAGES = $(TOKENS) $(UTILS)

all: build install

build: $(PACKAGES)
	go build $(UTILS)
	go build $(TOKENS)

install:
	go install
