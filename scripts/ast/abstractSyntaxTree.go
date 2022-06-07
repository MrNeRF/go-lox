package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var astMaps = map[string]map[string][]string{
	"Expr": expressionAst,
	"Stmt": statementAst,
}

// ast = Abstract Syntax tree
var expressionAst = map[string][]string{
	"Binary":   {"Left Expr", "Operator tokens.Token", "Right Expr"},
	"Grouping": {"Expression Expr"},
	"Literal":  {"Value interface{}"},
	"Unary":    {"Operator tokens.Token", "Right Expr"},
}

var statementAst = map[string][]string{
	"Expression": {"Expression Expr"},
	"Print":      {"Expression Expr"},
}

func CreateInterface(interfaceName string) string {
	tmpstring := fmt.Sprintf("type %v interface {\n", interfaceName)
	tmpstring += fmt.Sprintf("\tAccept(visitor %vVisitor) interface{}\n", interfaceName)
	tmpstring += "}\n"
	return tmpstring
}

func CreateConstructor(nameDerivedClass string) string {
	return fmt.Sprintf("func New%s() *%s {\n\treturn &%s{}\n}\n", nameDerivedClass, nameDerivedClass, nameDerivedClass)
}

func CreateClass(nameDerivedClass string, ast map[string][]string) string {
	sl := ast[nameDerivedClass]
	tmpString := fmt.Sprintf("type %v struct {\n", nameDerivedClass)

	for _, value := range sl {
		tmpString = tmpString + fmt.Sprintf("\t%s\n", value)
	}
	tmpString = tmpString + "}\n"
	return tmpString
}

func CreateVisitorInterface(prefix string, ast map[string][]string) string {
	tmpString := fmt.Sprintf("type %vVisitor interface {\n", prefix)
	for key := range ast {
		tmpString += fmt.Sprintf("\tvisit%s(e *%s) interface{}\n", key, key)
	}
	tmpString += "}\n"
	return tmpString
}

func CreateAcceptMethod(key string, interfaceName string) string {
	tmpString := fmt.Sprintf("func (e *%s) Accept(visitor %vVisitor) interface{} {\n", key, interfaceName)
	tmpString += fmt.Sprintf("\treturn visitor.visit%s(e)\n", key)
	tmpString += "}\n"
	return tmpString
}

func createAst(prefix string, ast map[string][]string) string {
	astString := CreateInterface(prefix) + "\n"
	for key := range ast {
		astString += CreateConstructor(key) + "\n"
		astString += CreateClass(key, ast) + "\n"
		astString += CreateAcceptMethod(key, prefix) + "\n"
	}
	astString += CreateVisitorInterface(prefix, ast) + "\n"
	return astString
}

func main() {
	file, err := os.Create("pkg/parser/ast.go")
	if err != nil {
		log.Fatal(err)
	}
	writer := bufio.NewWriter(file)
	astString := "package parser\n\n"
	astString += "import \"go-lox/pkg/tokens\"\n\n"
	for key, val := range astMaps {
		astString += createAst(key, val)
	}
	writer.WriteString(astString)
	writer.Flush()
}
