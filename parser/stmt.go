package parser

import (
	"fmt"
	"monkey-interpreter/ast"
	"monkey-interpreter/token"
)

func (p *Parser) parseStatement() (ast.Statement, error) {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLet()

	case token.NEWLINE:
		return p.parseEmpty()

	default:
		return nil, fmt.Errorf("Invalid token for statement: %v", p.curToken)
	}
}

func (p *Parser) parseEmpty() (ast.Statement, error) {
	stmt := ast.EmptyStatement{Token: p.curToken}

	err := p.nextToken()
	if err != nil {
		return nil, err
	}

	return &stmt, nil
}

func (p *Parser) parseLet() (ast.Statement, error) {
	// expect to only be called when curToken is let
	if p.curToken.Type != token.LET {
		panic(fmt.Sprintf("parseLet called with invalid token %v", p.curToken))
	}

	letStmt := ast.LetStatement{
		Token: p.curToken,
	}

	// identifier
	p.nextToken()

	if p.curToken.Type != token.IDENT {
		return &letStmt, fmt.Errorf("Expected identifier after let, got %v", p.curToken)
	}

	letStmt.Ident = ast.Identifier{
		Token: p.curToken,
		Name:  p.curToken.Literal,
	}

	// assign (=)
	p.nextToken()

	if p.curToken.Type != token.ASSIGN {
		return &letStmt, fmt.Errorf("Expected assignment (=), got %v", p.curToken)
	}

	// expression
	p.nextToken()

	expr, err := p.parseExpr()
	if err != nil {
		return &letStmt, err
	}

	letStmt.Expr = expr

	// newline
	p.nextToken()
	if p.curToken.Type != token.NEWLINE {
		return &letStmt, fmt.Errorf("Expected newline (\\n), got %v", p.curToken)
	}

	p.nextToken()

	return &letStmt, nil
}
