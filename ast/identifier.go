package ast

import "monkey-interpreter/token"

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
