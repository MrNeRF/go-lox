build: src/*.go
	go build -o build/go-lox $^

.PHONY: clean
clean:
	rm -rf build