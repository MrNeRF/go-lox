package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// ast = Abstract Syntax tree
var ast = map[string][]string{
	"Binary":   {"left Expr", "operator tokens.Token", "right Expr"},
	"Grouping": {"expression Expr"},
	"Literal":  {"value interface{}"},
	"Unary":    {"operator tokens.Token", "right Expr"},
}

func CreateExpression() string {
	return fmt.Sprintln("type Expr interface {}")
}

func CreateConstructor(nameDerivedClass string) string {
	return fmt.Sprintf("func New%s() *%s {\n\treturn &%s{}\n}\n", nameDerivedClass, nameDerivedClass, nameDerivedClass)
}

func CreateClass(nameDerivedClass string) string {
	sl := ast[nameDerivedClass]
	tmpString := fmt.Sprintf("type %v struct {\n", nameDerivedClass)

	for _, value := range sl {
		tmpString = tmpString + fmt.Sprintf("\t%s\n", value)
	}
	tmpString = tmpString + "}\n"
	return tmpString
}

func CreateVisitorInterface() string {
	tmpString := "type ExprVisitor interface {\n"
	for key := range ast {
		tmpString += fmt.Sprintf("\tvisit%s(e *%s) interface{}\n", key, key)
	}
	tmpString += "}\n"
	return tmpString
}

func CreateAcceptMethod(key string) string {
	tmpString := fmt.Sprintf("func (e *%s) Accept(visitor ExprVisitor) interface{} {\n", key)
	tmpString += fmt.Sprintf("\treturn visitor.visit%s(e)\n", key)
	tmpString += "}\n"
	return tmpString
}

func main() {
	file, err := os.Create("pkg/parser/ast.go")
	if err != nil {
		log.Fatal(err)
	}
	writer := bufio.NewWriter(file)
	astString := "package parser\n\n"
	astString += "import \"go-lox/pkg/tokens\"\n\n"
	astString += CreateExpression() + "\n"
	for key := range ast {
		astString += CreateConstructor(key) + "\n"
		astString += CreateClass(key) + "\n"
		astString += CreateAcceptMethod(key) + "\n"
	}
	astString += CreateVisitorInterface() + "\n"
	writer.WriteString(astString)
	writer.Flush()
}
