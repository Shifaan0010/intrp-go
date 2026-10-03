package parser

import (
	"errors"
	"fmt"
	"intrp-go/ast"
	"intrp-go/parser/precedence"
	"intrp-go/token"
)

func (p *Parser) parseExpr(prec precedence.Precedence) (ast.Expression, error) {
	// if !isPrefix(p.curToken) {
	//
	// }

	leftExpr, err := p.parsePrefix()
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

func (p *Parser) parsePrefix() (ast.Expression, error) {
	var leftExpr ast.Expression
	var err error

	switch p.curToken.Type {
	case token.INT:
		leftExpr, err = p.parseInt()

	case token.TRUE:
		fallthrough
	case token.FALSE:
		leftExpr, err = p.parseBool()

	case token.IDENT:
		leftExpr, err = p.parseIdent()

	case token.BANG:
		fallthrough
	case token.MINUS:
		leftExpr, err = p.parsePrefixExpr()

	case token.LPAREN:
		leftExpr, err = p.parseParenthesis()

	case token.LBRACE:
		leftExpr, err = p.parseBlock()

	case token.IF:
		leftExpr, err = p.parseIf()

	case token.FUNCTION:
		leftExpr, err = p.parseFnDecl()

	default:
		return nil, fmt.Errorf("no prefix fn for token %s", p.curToken.Type)
	}

	if err != nil {
		return nil, errors.Join(errors.New("parseExpr: failed to parse prefix"), err)
	}

	return leftExpr, nil
}

func (p *Parser) parsePrefixExpr() (ast.Expression, error) {
	var err error = nil

	expr := &ast.PrefixExpr{
		Op: p.curToken,
	}

	p.nextToken()

	switch p.curToken.Type {
	case token.INT:
		expr.Expr, err = p.parseInt()

	case token.TRUE:
		fallthrough
	case token.FALSE:
		expr.Expr, err = p.parseBool()

	case token.IDENT:
		expr.Expr, err = p.parseIdent()

	case token.LPAREN:
		expr.Expr, err = p.parsePrefix()

	default:
		err = fmt.Errorf("unexpected token %v for parsePrefixExpr", p.curToken)
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
		tok == token.GT ||
		tok == token.LTE ||
		tok == token.GTE ||
		tok == token.LPAREN ||
		tok == token.COMMA
}

func (p *Parser) parseInfix(left ast.Expression, prec precedence.Precedence) (ast.Expression, error) {
	switch p.curToken.Type {
	case token.LPAREN:
		expr := &ast.FnCallExpr{
			FnExpr: left,
			Tok: p.curToken,
		}
		argsExpr, err := p.parseFnArgs()
		if err != nil {
			return nil, err
		}

		expr.Params = argsExpr

		return expr, nil

	default:
		expr := &ast.InfixExpr{
			Op:   p.curToken,
			Left: left,
		}

		p.nextToken()

		rightExpr, err := p.parseExpr(prec)
		if err != nil {
			return expr, err
		}

		expr.Right = rightExpr

		return expr, nil
	}
}

func (p *Parser) parseParenthesis() (ast.Expression, error) {
	if p.curToken.Type != token.LPAREN {
		panic(fmt.Sprintf("parseParenthesis called with invalid token %s", p.curToken))
	}

	p.nextToken()

	leftExpr, err := p.parseExpr(precedence.LOWEST)
	if err != nil {
		return leftExpr, err
	}

	if p.curToken.Type != token.RPAREN {
		return leftExpr, fmt.Errorf("parseParenthesis: expected ')', got %s", p.curToken)
	}

	p.nextToken()

	return leftExpr, err
}

func (p *Parser) parseIf() (ast.Expression, error) {
	if p.curToken.Type != token.IF {
		panic(fmt.Sprintf("parseIf called with invalid token %s", p.curToken))
	}

	expr := &ast.IfExpr{
		Tok: p.curToken,
	}

	p.nextToken()

	var err error

	expr.Cond, err = p.parseExpr(precedence.LOWEST)
	if err != nil {
		return expr, err
	}

	expr.If, err = p.parseExpr(precedence.LOWEST)
	if err != nil {
		return expr, err
	}

	if p.curToken.Type == token.ELSE {
		p.nextToken()

		elseExpr, err := p.parseExpr(precedence.LOWEST)
		if err != nil {
			return expr, err
		}

		expr.Else = &elseExpr
	}

	return expr, nil
}

func (p *Parser) parseFnDecl() (ast.Expression, error) {
	if p.curToken.Type != token.FUNCTION {
		panic(fmt.Sprintf("parseFn called with invalid token %s", p.curToken))
	}

	expr := &ast.FnDeclExpr{
		Tok: p.curToken,
	}

	p.nextToken()
	if p.curToken.Type != token.LPAREN {
		return nil, fmt.Errorf("parseFn: expected (, got %s", p.curToken)
	}

	p.nextToken()

	expr.Params = []ast.Identifier{}
	for p.curToken.Type == token.IDENT {
		ident, _ := p.parseIdent()

		expr.Params = append(expr.Params, *ident)

		if p.curToken.Type == token.COMMA {
			p.nextToken()
			if p.curToken.Type != token.IDENT {
				return nil, fmt.Errorf("parseFn: expected Identifier after comma, got %s", p.curToken)
			}
		}
	}

	if p.curToken.Type != token.RPAREN {
		return nil, fmt.Errorf("parseFn: expected ), got %s", p.curToken)
	}

	p.nextToken()

	block, err := p.parseBlock()
	if err != nil {
		return expr, err
	}

	expr.Block = *block

	return expr, nil
}

func (p *Parser) parseFnArgs() ([]ast.Expression, error) {
	if p.curToken.Type != token.LPAREN {
		panic(fmt.Sprintf("parseFnArgs called with invalid token %s", p.curToken))
	}

	p.nextToken()

	if p.curToken.Type == token.RPAREN {
		p.nextToken()

		return []ast.Expression{}, nil
	}

	args := []ast.Expression{}
	for p.curToken.Type != token.RPAREN {
		argExpr, _ := p.parseExpr(precedence.COMMA)

		args = append(args, argExpr)

		if p.curToken.Type == token.COMMA {
			p.nextToken()
		}
	}

	if p.curToken.Type != token.RPAREN {
		return args, fmt.Errorf("parseFnArgs: expected ')', got %s", p.curToken)
	}

	p.nextToken()

	return args, nil
}

func (p *Parser) parseBlock() (*ast.BlockExpr, error) {
	if p.curToken.Type != token.LBRACE {
		panic(fmt.Sprintf("parseBlock called with invalid token %s", p.curToken))
	}

	block := &ast.BlockExpr{
		Stmts: []ast.Statement{},
	}

	p.nextToken()

	for p.curToken.Type != token.RBRACE {
		stmt, err := p.parseStatement()
		if err != nil {
			return block, err
		}

		block.Stmts = append(block.Stmts, stmt)
	}

	p.nextToken()

	return block, nil
}
