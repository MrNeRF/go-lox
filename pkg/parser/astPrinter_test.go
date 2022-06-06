package parser

import (
	"go-lox/pkg/tokens"
	"testing"
)

func TestPrint(t *testing.T) {

	var expr Expr = &Binary{
		Left: &Unary{
			Operator: *tokens.NewToken(tokens.MINUS, "-", 1, nil),
			Right:    &Literal{Value: 123}},
		Operator: *tokens.NewToken(tokens.STAR, "*", 1, nil),
		Right:    &Grouping{Expression: &Literal{Value: 45.67}}}

	astPrinter := NewAstPrinter()
	result := astPrinter.Print(expr)
	expected := "(* (- 123) (group 45.67))"

	if result != expected {
		t.Fatalf("AstPrinter.Print(),\n result: %v,\n expected: %v\n", result, expected)
	}
}
