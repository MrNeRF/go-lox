package main

import (
	"bufio"
	"fmt"
	"os"
)

// run
func run(s string) {}

func runFile(path string) {}

func runPrompt() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		} else {
			fmt.Println(line)
			fmt.Print("> ")
		}
	}
}

func main() {
	args := os.Args
	if len(args) > 1 {
		fmt.Println("Usage: go-lox [script]")
		return
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}
}
