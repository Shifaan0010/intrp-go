package ast

import (
	"fmt"
	"monkey-interpreter/token"
	"strings"
)

// type Operator int
//
// const (
// 	UNKNOWN = iota
//
// 	BOOL_NOT
// 	PLUS
// 	MINUS
// 	MULT
// 	DIV
// 	GT
// 	LT
// )
//
// func (op Operator) String() string {
// 	switch op {
// 	case BOOL_NOT:
// 		return "!"
// 	case PLUS:
// 		return "+"
// 	case MINUS:
// 		return "-"
// 	case MULT:
// 		return "*"
// 	case DIV:
// 		return "/"
// 	case GT:
// 		return ">"
// 	case LT:
// 		return "<"
// 	}
//
// 	return "UNKNOWN"
// }

type PrefixExpr struct {
	Op   token.Token
	Expr Expression
}

func (e *PrefixExpr) TokenLiteral() string {
	return e.Op.String()
}

func (e *PrefixExpr) String() string {
	return fmt.Sprintf("(%s %s)", e.Op, e.Expr)
}

func (e *PrefixExpr) expressionNode() {}

type InfixExpr struct {
	Op    token.Token
	Left  Expression
	Right Expression
}

func (e *InfixExpr) TokenLiteral() string {
	return e.Op.Literal
}

func (e *InfixExpr) String() string {
	return fmt.Sprintf("(%s %s %s)", e.Left, e.Op, e.Right)
}

func (e *InfixExpr) expressionNode() {}

type IfExpr struct {
	Cond Expression
	If   Expression
	Else *Expression
	Tok  token.Token
}

func (e *IfExpr) TokenLiteral() string {
	return e.Tok.Literal
}

func (e *IfExpr) String() string {
	if e.Else == nil {
		return fmt.Sprintf("if %s %s", e.Cond, e.If)
	} else {
		return fmt.Sprintf("if %s %s else %s", e.Cond, e.If, *e.Else)
	}
}

func (e *IfExpr) expressionNode() {}

type FnExpr struct {
	Params []Identifier
	Block  BlockExpr
	Tok    token.Token
}

func (e *FnExpr) TokenLiteral() string {
	return e.Tok.Literal
}

func (e *FnExpr) String() string {
	return fmt.Sprintf("fn (%s) %s", &e.Params, e.Block)
}

func (e *FnExpr) expressionNode() {}

type BlockExpr struct {
	Stmts []Statement
	Tok   token.Token
}

func (e *BlockExpr) TokenLiteral() string {
	return e.Tok.Literal
}

func (e *BlockExpr) String() string {
	sb := strings.Builder{}

	sb.WriteString("{")

	for _, stmt := range e.Stmts {
		sb.WriteString(stmt.String())
	}

	sb.WriteString("}")

	return sb.String()
}

func (e *BlockExpr) expressionNode() {}
