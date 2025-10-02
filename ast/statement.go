package ast

import "monkey-interpreter/token"

type LetStatement struct {
	Token token.Token
	Ident Identifier
	Expr  Expression
}

func (p *LetStatement) TokenLiteral() string {
	return p.Token.Literal
}

func (l *LetStatement) statementNode() {}

type EmptyStatement struct {
	Token token.Token
}

func (p *EmptyStatement) TokenLiteral() string {
	return p.Token.Literal
}

func (l *EmptyStatement) statementNode() {}

