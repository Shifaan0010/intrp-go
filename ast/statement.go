package ast

import (
	"fmt"
	"intrp-go/token"
)

type LetStatement struct {
	Token  token.Token
	Assign AssignStatement
}

func (p *LetStatement) TokenLiteral() string {
	return p.Token.Literal
}

func (p *LetStatement) String() string {
	return fmt.Sprintf("let %s = %s\n", &p.Assign.Ident, p.Assign.Expr)
}

func (l *LetStatement) statementNode() {}

type ReturnStatement struct {
	Token token.Token
	Expr  Expression
}

func (p *ReturnStatement) TokenLiteral() string {
	return p.Token.Literal
}

func (p *ReturnStatement) String() string {
	return fmt.Sprintf("return %s\n", p.Expr)
}

func (l *ReturnStatement) statementNode() {}

type AssignStatement struct {
	Token token.Token
	Ident Identifier
	Expr  Expression
}

func (p *AssignStatement) TokenLiteral() string {
	return p.Token.Literal
}

func (p *AssignStatement) String() string {
	return fmt.Sprintf("%s = %s\n", &p.Ident, p.Expr)
}

func (l *AssignStatement) statementNode() {}

type EmptyStatement struct {
	Token token.Token
}

func (p *EmptyStatement) TokenLiteral() string {
	return p.Token.Literal
}

func (p *EmptyStatement) String() string {
	return "\n"
}

func (l *EmptyStatement) statementNode() {}

type ExprStatement struct {
	Token token.Token
	Expr  Expression
}

func (p *ExprStatement) TokenLiteral() string {
	return p.Token.Literal
}

func (p *ExprStatement) String() string {
	return p.Expr.String() + "\n"
}

func (l *ExprStatement) statementNode() {}
