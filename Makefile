#https://go.dev/doc/code

TOKENS = pkg/tokens/*.go
UTILS = pkg/utils/*.go
PARSER = pkg/parser/*.go
AST = scripts/ast/abstractSyntaxTree.go
PACKAGES = $(TOKENS) $(UTILS) $(PARSER)

all: ast build install

ast: $(AST)
	go build -o $(basename $(AST)) $(AST)
	./$(basename $(AST))
	rm $(basename $(AST))

build: $(PACKAGES)
	go build $(UTILS)
	go build $(TOKENS)
	go build $(PARSER)

install:
	go install

test:
	go test -cover -v ./pkg/parser/...