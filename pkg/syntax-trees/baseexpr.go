package syntaxtrees

import (
	"fmt"
	"go-lox/pkg/tokens"
)

type Expr interface {
}

type Binary struct {
	left     Expr
	operator tokens.Token
	right    Expr
}

type Unary struct {
	left     Expr
	operator tokens.Token
}

func (b *Binary) Accept(visitor ExprVisitor) interface{} {
	return visitor.visitBinary(b)
}

func (u *Unary) Accept(visitor ExprVisitor) interface{} {
	return visitor.visitUnary(u)
}

type ExprVisitor interface {
	visitBinary(b *Binary) interface{}
	visitUnary(u *Unary) interface{}
	//...
}

//Lets do the fuck visitor
type FuckVisitor struct {
	times int
}

func NewFuckVisitor(times int) *FuckVisitor {
	return &FuckVisitor{times: times}
}

func (fck *FuckVisitor) visitBinary(b *Binary) interface{} {
	return "a fucked binary!"
}

func (fck *FuckVisitor) visitUnary(u *Unary) interface{} {
	fmt.Println("I can also fuck Unarys and return how many times I fucked him/her/it/+!")
	return fck.times
}
