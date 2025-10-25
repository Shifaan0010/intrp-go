package ast

import (
	"fmt"
	"monkey-interpreter/token"
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

// type AddExpr struct {
// 	Token token.Token
// 	left  Expression
// 	right Expression
// }
//
// func (i *AddExpr) TokenLiteral() string {
// 	return i.Token.Literal
// }
//
// func (i *AddExpr) String() string {
// 	return fmt.Sprintf("%s + %s", i.left, i.right)
// }
//
// func (i *AddExpr) expressionNode() {}
