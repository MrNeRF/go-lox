#https://go.dev/doc/code

MAIN = src/main.go
TOKENS = src/tokens/*.go

#build: $(SOURCES)
#	go build -o build/go-lox $^

build:
	go build $(TOKENS)
	go build -o build/go-lox


.PHONY: clean
clean:
	rm -rf build