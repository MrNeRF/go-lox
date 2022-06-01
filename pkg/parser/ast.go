package parser

import "go-lox/pkg/tokens"

type Expr interface {}

func NewGrouping() *Grouping {
	return &Grouping{}
}

type Grouping struct {
	expression Expr
}

func (e *Grouping) Accept(visitor ExprVisitor) interface{} {
	return visitor.visitGrouping(e)
}

func NewLiteral() *Literal {
	return &Literal{}
}

type Literal struct {
	value interface{}
}

func (e *Literal) Accept(visitor ExprVisitor) interface{} {
	return visitor.visitLiteral(e)
}

func NewUnary() *Unary {
	return &Unary{}
}

type Unary struct {
	operator tokens.Token
	right Expr
}

func (e *Unary) Accept(visitor ExprVisitor) interface{} {
	return visitor.visitUnary(e)
}

func NewBinary() *Binary {
	return &Binary{}
}

type Binary struct {
	left Expr
	operator tokens.Token
	right Expr
}

func (e *Binary) Accept(visitor ExprVisitor) interface{} {
	return visitor.visitBinary(e)
}

type ExprVisitor interface {
	visitBinary(e *Binary) interface{}
	visitGrouping(e *Grouping) interface{}
	visitLiteral(e *Literal) interface{}
	visitUnary(e *Unary) interface{}
}

