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
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
		fmt.Print("> ")
	}
}

func main() {
	runPrompt()
}
