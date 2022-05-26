package main

import (
	"bufio"
	"fmt"
	"os"
)

// run runs a line of lox
func run(s string) {
	fmt.Println(s)
}

func runFile(path string) {}

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
