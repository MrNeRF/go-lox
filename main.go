package main

import (
	"bufio"
	"fmt"
	"os"
)

func runPrompt() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		} else {
			fmt.Println(line)
			fmt.Print("> ")
		}
	}
}

func main() {
	runPrompt()
}
