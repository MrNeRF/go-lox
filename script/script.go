package main

import "fmt"

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
	fmt.Println("package syntaxtrees")
	fmt.Println("import \"go-lox/pkg/tokens\"")
	fmt.Println(CreateExpression())
	for key := range ast {
		fmt.Println(CreateConstructor(key))
		fmt.Println(CreateClass(key))
		fmt.Println(CreateAcceptMethod(key))
	}
	fmt.Println(CreateVisitorInterface())
}
