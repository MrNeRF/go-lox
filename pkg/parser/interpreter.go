package parser

import "go-lox/pkg/tokens"

// implements ExprVisitor
type Interpreter struct {
}

func (ip *Interpreter) visitLiteral(e *Literal) interface{} {
	return e.Value
}

func (ip *Interpreter) visitGrouping(e *Grouping) interface{} {
	return ip.evaluate(e.Expression)
}

func (ip *Interpreter) visitUnary(e *Unary) interface{} {
	var right interface{} = ip.evaluate(e.Right)

	switch e.Operator.GetTokenType() {
	case tokens.MINUS:
		return -right.(float64)
	}
	return nil
}

func (ip *Interpreter) visitBinary(e *Binary) interface{} {
	return ip.evaluate(e.Left)
}

func (ip *Interpreter) evaluate(expr Expr) interface{} {
	return expr.Accept(ip)
}
