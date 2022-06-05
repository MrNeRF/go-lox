package parser

import "go-lox/pkg/tokens"

type Expr interface {
	Accept(visitor ExprVisitor) interface{}
}

func NewBinary() *Binary {
	return &Binary{}
}

type Binary struct {
	Left Expr
	Operator tokens.Token
	Right Expr
}

func (e *Binary) Accept(visitor ExprVisitor) interface{} {
	return visitor.visitBinary(e)
}

func NewGrouping() *Grouping {
	return &Grouping{}
}

type Grouping struct {
	Expression Expr
}

func (e *Grouping) Accept(visitor ExprVisitor) interface{} {
	return visitor.visitGrouping(e)
}

func NewLiteral() *Literal {
	return &Literal{}
}

type Literal struct {
	Value interface{}
}

func (e *Literal) Accept(visitor ExprVisitor) interface{} {
	return visitor.visitLiteral(e)
}

func NewUnary() *Unary {
	return &Unary{}
}

type Unary struct {
	Operator tokens.Token
	Right Expr
}

func (e *Unary) Accept(visitor ExprVisitor) interface{} {
	return visitor.visitUnary(e)
}

type ExprVisitor interface {
	visitBinary(e *Binary) interface{}
	visitGrouping(e *Grouping) interface{}
	visitLiteral(e *Literal) interface{}
	visitUnary(e *Unary) interface{}
}

