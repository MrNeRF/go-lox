package parser

import (
	"errors"
	"fmt"
	"go-lox/pkg/tokens"
	"log"
)

// implements ExprVisitor
type Interpreter struct {
}

func NewInterpreter() *Interpreter {
	return &Interpreter{}
}

func (ip *Interpreter) Interpret(expr Expr) {
	value := ip.evaluate(expr)
	fmt.Println(stringify(value))
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
		err := checkNumberOperand(e.Operator, right)
		if err != nil {
			log.Println(err.Error()) //Here we need better Error handling! Maybe a special class for Errors.
		}
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
		err := checkNumberOperands(e.Operator, left, right)
		if err != nil {
			log.Println(err.Error()) //Here we need better Error handling! Maybe a special class for Errors.
		}
		return left.(float64) > right.(float64)
	case tokens.GREATER_EQUAL:
		err := checkNumberOperands(e.Operator, left, right)
		if err != nil {
			log.Println(err.Error()) //Here we need better Error handling! Maybe a special class for Errors.
		}
		return left.(float64) >= right.(float64)
	case tokens.LESS:
		err := checkNumberOperands(e.Operator, left, right)
		if err != nil {
			log.Println(err.Error()) //Here we need better Error handling! Maybe a special class for Errors.
		}
		return left.(float64) < right.(float64)
	case tokens.LESS_EQUAL:
		err := checkNumberOperands(e.Operator, left, right)
		if err != nil {
			log.Println(err.Error()) //Here we need better Error handling! Maybe a special class for Errors.
		}
		return left.(float64) <= right.(float64)
	case tokens.BANG_EQUAL:
		return !isEqual(left, right)
	case tokens.EQUAL_EQUAL:
		return isEqual(left, right)
	case tokens.MINUS:
		err := checkNumberOperands(e.Operator, left, right)
		if err != nil {
			log.Println(err.Error()) //Here we need better Error handling! Maybe a special class for Errors.
		}
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
		log.Println("Operands must be two numbers or two strings") //Here we need a better Error handling!

	case tokens.SLASH:
		err := checkNumberOperands(e.Operator, left, right)
		if err != nil {
			log.Println(err.Error()) //Here we need better Error handling! Maybe a special class for Errors.
		}
		return left.(float64) / right.(float64)
	case tokens.STAR:
		err := checkNumberOperands(e.Operator, left, right)
		if err != nil {
			log.Println(err.Error()) //Here we need better Error handling! Maybe a special class for Errors.
		}
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

func checkNumberOperand(tk tokens.Token, operand interface{}) error {
	if _, ok := operand.(float64); ok {
		return nil
	} else {
		errStr := fmt.Sprintf("Operand %v must be a number", operand)
		return errors.New(errStr)
	}
}

func checkNumberOperands(tk tokens.Token, leftOperand interface{}, rightOperand interface{}) error {
	_, leftokfloat := leftOperand.(float64)
	_, rightokfloat := rightOperand.(float64)
	if leftokfloat && rightokfloat {
		return nil
	}
	var errStr string = ""
	if !leftokfloat {
		errStr += fmt.Sprintf("Left Operand %v must be a number\n", leftOperand)
	}
	if !rightokfloat {
		errStr += fmt.Sprintf("Right Operand %v must be a number\n", rightOperand)
	}
	return errors.New(errStr)
}

func stringify(expr interface{}) string {

	if expr == nil {
		return "nil"
	}
	if val, ok := expr.(float64); ok {
		str := fmt.Sprintf("%v", val)
		return str
	}
	if val, ok := expr.(bool); ok {
		str := fmt.Sprintf("%v", val)
		return str
	}
	val, ok := expr.(string)
	if !ok {
		log.Panic("stringify() could not be successfully executed")
	}
	return val
}
