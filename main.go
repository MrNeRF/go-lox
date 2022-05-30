package main

import (
	"bufio"
	"fmt"
	syntaxtrees "go-lox/pkg/syntax-trees"
	"go-lox/pkg/tokens"
	"go-lox/pkg/utils"
	"io/ioutil"
	"os"
)

// run runs a line of lox
func run(s string) {
	t := tokens.NewTokenizer(s)
	t.ScanTokens()
	fmt.Println(t.GetTokenList())
}

func runFile(path string) {
	content, err := ioutil.ReadFile(path) // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	run(string(content))
}

func runPrompt() {

	utils.PrintPromptInit()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">>> ")
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
	fucker := syntaxtrees.NewFuckVisitor(1)
	apoorUnary := syntaxtrees.Unary{}
	//and now...we fuck it
	apoorUnary.Accept(fucker)

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
