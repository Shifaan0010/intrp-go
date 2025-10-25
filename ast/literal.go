package ast

import (
	"monkey-interpreter/token"
	"strconv"
)

type IntLiteral struct {
	Token token.Token
	Val   int64
}

func (e *IntLiteral) TokenLiteral() string {
	return e.Token.Literal
}

func (e *IntLiteral) String() string {
	return strconv.FormatInt(e.Val, 10)
}

func (e *IntLiteral) expressionNode() {}

type BoolLiteral struct {
	Token token.Token
	Val   bool
}

func (e *BoolLiteral) TokenLiteral() string {
	return e.Token.Literal
}

func (e *BoolLiteral) String() string {
	return strconv.FormatBool(e.Val)
}

func (e *BoolLiteral) expressionNode() {}

type Identifier struct {
	Token token.Token
	Name  string
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Token.Literal
}

func (i *Identifier) expressionNode() {}
