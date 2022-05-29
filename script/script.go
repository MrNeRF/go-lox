package main

import "fmt"

// ast = Abstract Syntax tree
var ast = map[string][]string{
	"Binary":   {"left Expr", "operator tokens.Token", "right Expr"},
	"Grouping": {"expression Expr"},
	"Literal":  {"value Interface{}"},
	"Unary":    {"operator Token", "right Expr"},
}

func CreateConstructor(nameDerivedClass string) string {
	return fmt.Sprintf("func New%s() *%s {\n\treturn &%s{}\n}", nameDerivedClass, nameDerivedClass, nameDerivedClass)
}

func CreateClass(nameDerivedClass string) string {
	sl := ast[nameDerivedClass]
	tmpString := fmt.Sprintf("type %v struct {\n", nameDerivedClass)

	for _, value := range sl {
		tmpString = tmpString + fmt.Sprintf("\t%s\n", value)
	}
	tmpString = tmpString + "}"
	return tmpString
}

func main() {

	fmt.Println(CreateConstructor("Binary"))
	fmt.Println(CreateClass("Binary"))
}
