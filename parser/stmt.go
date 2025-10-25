package parser

import (
	"errors"
	"fmt"
	"monkey-interpreter/ast"
	"monkey-interpreter/parser/precedence"
	"monkey-interpreter/token"
)

func (p *Parser) parseStatement() (ast.Statement, error) {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLet()

	case token.NEWLINE:
		return p.parseEmpty()

	default:
		return p.parseExprStmt()
		// return nil, fmt.Errorf("Invalid token for statement: %v", p.curToken)
	}
}

func (p *Parser) parseEmpty() (ast.Statement, error) {
	// expect to only be called when curToken is newline
	if p.curToken.Type != token.NEWLINE {
		panic(fmt.Sprintf("parseEmpty called with invalid token %v", p.curToken))
	}

	stmt := ast.EmptyStatement{Token: p.curToken}

	err := p.nextToken()
	if err != nil {
		return nil, err
	}

	return &stmt, nil
}

func (p *Parser) parseExprStmt() (ast.Statement, error) {
	expr, err := p.parseExpr(precedence.LOWEST)

	if err != nil {
		return nil, errors.Join(fmt.Errorf("error while parsing expr"), err)
	}

	if isEndOfStmt(p.curToken.Type) {
		return nil, fmt.Errorf("Expected newline, got %v, %v", p.curToken, expr)
	}

	return &ast.ExprStatement{
		Token: token.Token{},
		Expr:  expr,
	}, err
}

func (p *Parser) parseLet() (ast.Statement, error) {
	// expect to only be called when curToken is let
	if p.curToken.Type != token.LET {
		panic(fmt.Sprintf("parseLet called with invalid token %v", p.curToken))
	}

	letStmt := ast.LetStatement{
		Token: p.curToken,
	}

	p.nextToken()

	// identifier
	if p.curToken.Type != token.IDENT {
		return &letStmt, fmt.Errorf("Expected identifier after let, got %v", p.curToken)
	}

	letStmt.Ident = ast.Identifier{
		Token: p.curToken,
		Name:  p.curToken.Literal,
	}

	p.nextToken()

	// assign (=)
	if p.curToken.Type != token.ASSIGN {
		return &letStmt, fmt.Errorf("Expected assignment (=), got %v", p.curToken)
	}

	p.nextToken()

	// expression
	expr, err := p.parseExpr(precedence.LOWEST)
	if err != nil {
		return &letStmt, err
	}

	letStmt.Expr = expr

	// newline
	if isEndOfStmt(p.curToken.Type) {
		return &letStmt, fmt.Errorf("Expected newline (\\n), got %v", p.curToken)
	}

	p.nextToken()

	return &letStmt, nil
}

func isEndOfStmt(tok token.TokenType) bool {
	return tok != token.NEWLINE && tok != token.EOF
}
