package utils

import "fmt"

const GoloxVersion = "0.0.1"
const GoVersion = "1.18"

func PrintPromptInit() {
	s := fmt.Sprintf("Go-Lox %v [go v%v] on Ubuntu 22.04 LTS", GoloxVersion, GoVersion)
	fmt.Println(s)
	fmt.Println("Type \"help\", \"copyright\", \"credits\" or \"license\" for more information.")
}
