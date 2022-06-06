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
	case tokens.BANG:
		return !isTruthy(right)
	}

	return nil
}

func (ip *Interpreter) visitBinary(e *Binary) interface{} {
	var left interface{} = ip.evaluate(e.Left)
	var right interface{} = ip.evaluate(e.Right)

	switch e.Operator.GetTokenType() {
	case tokens.GREATER:
		return left.(float64) > right.(float64)
	case tokens.GREATER_EQUAL:
		return left.(float64) >= right.(float64)
	case tokens.LESS:
		return left.(float64) < right.(float64)
	case tokens.LESS_EQUAL:
		return left.(float64) <= right.(float64)
	case tokens.BANG_EQUAL:
		return !isEqual(left, right)
	case tokens.EQUAL_EQUAL:
		return isEqual(left, right)
	case tokens.MINUS:
		return left.(float64) - right.(float64)
	case tokens.PLUS:
		leftfloat, leftokfloat := left.(float64)
		rightfloat, rightokfloat := right.(float64)
		if leftokfloat && rightokfloat {
			return leftfloat + rightfloat
		}

		leftstring, leftokstring := left.(string)
		rightstring, rightokstring := right.(string)
		if leftokstring && rightokstring {
			return leftstring + rightstring
		}
	case tokens.SLASH:
		return left.(float64) / right.(float64)
	case tokens.STAR:
		return left.(float64) * right.(float64)
	}

	return nil
}

func (ip *Interpreter) evaluate(expr Expr) interface{} {
	return expr.Accept(ip)
}

func isTruthy(expr interface{}) bool {
	if expr == nil {
		return false
	}
	if bval, ok := expr.(bool); ok {
		return bval
	}
	return true
}

func isEqual(a interface{}, b interface{}) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return false
	}
	return a == b
}
