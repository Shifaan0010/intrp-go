package eval

import (
	"bufio"
	"fmt"
	"intrp-go/ast"
	"intrp-go/lexer"
	"intrp-go/object"
	"intrp-go/parser"
	"intrp-go/token"
	"strings"
)

func (e *Environment) Eval(stmtStr string) string {
	l := lexer.New(*bufio.NewReader(strings.NewReader(stmtStr)))

	p, err := parser.New(l)
	if err != nil {
		return fmt.Sprintf("failed to init parser, err %s", err)
	}

	if p.IsAtEof() {
		return ""
	}

	stmt, err := p.ParseStatement()
	if err != nil {
		return fmt.Sprintf("failed to parse statement, err %s", err)
	}

	// fmt.Printf("parsed statement: %#v\n", stmt)
	// fmt.Printf("parsed statement: %s\n", stmt)

	evald, err := e.EvalNode(stmt)
	if err != nil {
		return fmt.Sprintf("error: %s", err)
	}

	return fmt.Sprint(evald)
}

func (e *Environment) EvalNode(node ast.Node) (object.Object, error) {
	switch t := node.(type) {
	case *ast.ExprStatement:
		return e.evalExpr(t.Expr)

	default:
		return nil, fmt.Errorf("EvalNode not implemented for node %s", node.String())
	}
}

func (e *Environment) evalExpr(node ast.Expression) (object.Object, error) {
	switch t := node.(type) {
	case *ast.IntLiteral:
		return &object.Integer{Val: t.Val}, nil

	case *ast.BoolLiteral:
		return &object.Boolean{Val: t.Val}, nil

	case *ast.Identifier:
		return nil, nil

	case *ast.InfixExpr:
		return e.evalInfix(t)

	case *ast.PrefixExpr:
		return e.evalPrefix(t)

	default:
		return nil, fmt.Errorf("evalExpr not implemented for node %s", node.String())
	}
}

func (e *Environment) evalInfix(expr *ast.InfixExpr) (object.Object, error) {
	left, lErr := e.evalExpr(expr.Left)
	if lErr != nil {
		return nil, lErr
	}

	right, rErr := e.evalExpr(expr.Right)
	if rErr != nil {
		return nil, rErr
	}

	switch expr.Op.Type {
	case token.PLUS:
		return object.Add(left, right)

	case token.MINUS:
		return object.Sub(left, right)

	case token.ASTERISK:
		return object.Mult(left, right)

	case token.SLASH:
		return object.Div(left, right)

	default:
		return nil, fmt.Errorf("evalInfix not implemented for op %s", expr.Op.Type)
	}
}

func (e *Environment) evalPrefix(expr *ast.PrefixExpr) (object.Object, error) {
	right, rErr := e.evalExpr(expr.Expr)
	if rErr != nil {
		return nil, rErr
	}

	switch expr.Op.Type {
	case token.MINUS:
		return object.Neg(right)

	case token.BANG:
		return object.Not(right)

	default:
		return nil, fmt.Errorf("evalPrefix not implemented for op %s", expr.Op.Type)
	}
}
