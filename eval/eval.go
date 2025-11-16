package eval

import (
	"fmt"
	"intrp-go/ast"
	"intrp-go/object"
	"intrp-go/token"
)

func Eval(node ast.Node) (object.Object, error) {
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
