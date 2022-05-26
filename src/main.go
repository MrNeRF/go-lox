package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// run runs a line of lox
func run(s string) {
	t := NewTokenizer(s)
	t.scanTokens()
}

func runFile(path string) {
	content, err := ioutil.ReadFile(path) // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	t := NewTokenizer(string(content))
	t.scanToken()
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
