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

func Eval(stmtStr string) string {
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

	evald, err := EvalNode(stmt)
	if err != nil {
		return fmt.Sprintf("error: %s", err)
	}

	return fmt.Sprint(evald)
}

func EvalNode(node ast.Node) (object.Object, error) {
	switch t := node.(type) {
	case *ast.ExprStatement:
		return evalExpr(t.Expr)

	default:
		return nil, nil
	}
}

func evalExpr(node ast.Expression) (object.Object, error) {
	switch t := node.(type) {
	case *ast.IntLiteral:
		return &object.Integer{Val: t.Val}, nil

	case *ast.BoolLiteral:
		return &object.Boolean{Val: t.Val}, nil

	case *ast.InfixExpr:
		return evalInfix(t)

	// case *ast.PrefixExpr:
	// 	return evalPrefix(t)

	default:
		return nil, fmt.Errorf("evalExpr not implemented for node %s", node.String())
	}
}

func evalInfix(expr *ast.InfixExpr) (object.Object, error) {
	left, lErr := evalExpr(expr.Left)
	if lErr != nil {
		return nil, lErr
	}

	right, rErr := evalExpr(expr.Right)
	if rErr != nil {
		return nil, rErr
	}

	switch expr.Op.Type {
	case token.PLUS:
		return object.Add(left, right)

	default:
		return nil, fmt.Errorf("evalInfix not implemented for op %s", expr.Op.Type)
	}
}
