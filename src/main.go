package main

import (
	"bufio"
	"fmt"
	"go-lox/src/tokens"
	"io/ioutil"
	"os"
)

// run runs a line of lox
func run(s string) {
	t := tokens.NewTokenizer(s)
	t.ScanTokens()
}

func runFile(path string) {
	content, err := ioutil.ReadFile(path) // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	t := tokens.NewTokenizer(string(content))
	t.ScanTokens()
	fmt.Println(t.GetTokenList())
}

func runPrompt() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		line := scanner.Text()
		if line == "" {
			continue
		} else {
			run(line)
		}
	}
}

func main() {
	args := os.Args
	if len(args) > 2 {
		fmt.Println("Usage: go-lox [script]")
		return
	} else if len(args) == 2 {
		runFile(args[1])
	} else {
		runPrompt()
	}
}
