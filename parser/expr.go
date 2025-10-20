package parser

import (
	"errors"
	"fmt"
	"monkey-interpreter/ast"
	"monkey-interpreter/parser/precedence"
	"monkey-interpreter/token"
	"strconv"
)

func (p *Parser) parseExpr(prec precedence.Precedence) (ast.Expression, error) {
	// if !isPrefix(p.curToken) {
	// 	
	// }

	leftExpr, err := p.parsePrefix(prec)
	if err != nil {
		return nil, errors.Join(errors.New("parseExpr: failed to parse prefix"), err)
	}

	// fmt.Println("parsed prefix", leftExpr, p, prec, p.curToken.Type, precedence.TokenPrecedence(p.curToken.Type))

	for p.curToken.Type != token.NEWLINE && prec < precedence.TokenPrecedence(p.curToken.Type) {
		// fmt.Println(p.curToken)
		if !isInfix(p.curToken.Type) {
			return leftExpr, nil
		}

		leftExpr, err = p.parseInfix(leftExpr, precedence.TokenPrecedence(p.curToken.Type))

		if err != nil {
			return nil, errors.Join(errors.New("parseExpr: failed to parse infix"), err)
		}
	}

	return leftExpr, nil
}

func isPrefix(tok token.TokenType) bool {
	return tok == token.INT ||
		tok == token.BANG ||
		tok == token.MINUS
}

func (p *Parser) parsePrefix(prec precedence.Precedence) (ast.Expression, error) {
	var leftExpr ast.Expression
	var err error

	switch p.curToken.Type {
	case token.INT:
		leftExpr, err = p.parseInt()

	case token.BANG:
		fallthrough
	case token.MINUS:
		leftExpr, err = p.parsePrefixExpr(precedence.PREFIX)

	default:
		return nil, fmt.Errorf("no prefix fn for token %s", p.curToken.Type)
	}

	if err != nil {
		return nil, errors.Join(errors.New("parseExpr: failed to parse prefix"), err)
	}

	return leftExpr, nil
}

func (p *Parser) parsePrefixExpr(prec precedence.Precedence) (ast.Expression, error) {
	var err error = nil

	expr := &ast.PrefixExpr{
		Op: p.curToken,
	}

	switch p.peekToken.Type {
	case token.INT:
		p.nextToken()

		expr.Expr, err = p.parseInt()

	// case token.IDENT:
	// 	expr.Expr, err = p.parseIdentifier()

	default:
		err = fmt.Errorf("unexpected token %v for parsePrefix", p.peekToken)
	}

	return expr, err
}

func isInfix(tok token.TokenType) bool {
	return tok == token.EQ ||
		tok == token.NOT_EQ ||
		tok == token.PLUS ||
		tok == token.MINUS ||
		tok == token.ASTERISK ||
		tok == token.SLASH ||
		tok == token.LT ||
		tok == token.GT
}

func (p *Parser) parseInfix(left ast.Expression, prec precedence.Precedence) (ast.Expression, error) {
	expr := &ast.InfixExpr{
		Op:   p.curToken,
		Left: left,
	}

	p.nextToken()

	// prec := precedence.TokenPrecedence(p.curToken())

	var err error

	expr.Right, err = p.parseExpr(prec)

	return expr, err
}

func (p *Parser) parseInt() (ast.Expression, error) {
	// expect to only be called when curToken is int
	if p.curToken.Type != token.INT {
		panic(fmt.Sprintf("parseInt called with invalid token %v", p.curToken))
	}

	val, err := strconv.ParseInt(p.curToken.Literal, 10, 64)

	node := ast.IntLiteral{
		Val:   val,
		Token: p.curToken,
	}

	p.nextToken()

	return &node, err
}
