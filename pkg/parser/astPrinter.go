package parser

import (
	"fmt"
	"strings"
)

//implements ExprVisitor
type AstPrinter struct {
}

func NewAstPrinter() *AstPrinter {
	return &AstPrinter{}
}

func (ap *AstPrinter) Print(expr Expr) string {
	return fmt.Sprintf("%v", expr.Accept(ap))
}

func (ap *AstPrinter) visitGrouping(e *Grouping) interface{} {
	return ap.parenthesize("group", e.Expression)
}

func (ap *AstPrinter) visitLiteral(e *Literal) interface{} {
	if e.Value == nil {
		return "nil"
	}
	str := fmt.Sprintf("%v", e.Value)
	return str
}

func (ap *AstPrinter) visitUnary(e *Unary) interface{} {
	return ap.parenthesize(e.Operator.GetLexeme(), e.Right)
}

func (ap *AstPrinter) visitBinary(e *Binary) interface{} {
	return ap.parenthesize(e.Operator.GetLexeme(), e.Left, e.Right)
}

func (ap *AstPrinter) parenthesize(name string, exprs ...Expr) string {
	stringBuilder := strings.Builder{}
	stringBuilder.WriteString("(" + name)
	for _, expr := range exprs {
		stringBuilder.WriteString(" ")
		str := fmt.Sprintf("%v", expr.Accept(ap))
		stringBuilder.WriteString(str)
	}

	stringBuilder.WriteString(")")
	return stringBuilder.String()
}
