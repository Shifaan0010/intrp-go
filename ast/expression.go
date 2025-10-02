package ast

import "monkey-interpreter/token"

type IntLiteral struct {
	Token token.Token
	Val   int
}

func (i *IntLiteral) TokenLiteral() string {
	return i.Token.Literal
}

func (i *IntLiteral) expressionNode() {}

type AddExpr struct {
	Token token.Token
	left  Expression
	right Expression
}

func (i *AddExpr) TokenLiteral() string {
	return i.Token.Literal
}

func (i *AddExpr) expressionNode() {}
