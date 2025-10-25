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

	case token.RETURN:
		return p.parseReturn()

	case token.NEWLINE:
		return p.parseEmpty()

	default:
		if p.curToken.Type == token.IDENT && p.peekToken.Type == token.ASSIGN {
			return p.parseAssign()
		}

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

	if !isEndOfStmt(p.curToken.Type) {
		return nil, fmt.Errorf("Expected newline, got %v, %v", p.curToken, expr)
	}

	p.nextToken()

	return &ast.ExprStatement{
		Token: token.Token{},
		Expr:  expr,
	}, err
}

func (p *Parser) parseReturn() (ast.Statement, error) {
	// expect to only be called when curToken is return
	if p.curToken.Type != token.RETURN {
		panic(fmt.Sprintf("parseReturn called with invalid token %v", p.curToken))
	}

	stmt := &ast.ReturnStatement{
		Token: p.curToken,
	}

	p.nextToken()

	expr, err := p.parseExpr(precedence.LOWEST)
	if err != nil {
		return stmt, err
	}

	stmt.Expr = expr

	if !isEndOfStmt(p.curToken.Type) {
		return nil, fmt.Errorf("Expected newline, got %v", p.curToken)
	}

	p.nextToken()

	return stmt, nil
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

	assignStmt, err := p.parseAssign()
	letStmt.Assign = *assignStmt

	return &letStmt, err
}

func (p *Parser) parseAssign() (*ast.AssignStatement, error) {
	stmt := &ast.AssignStatement{}

	// identifier
	if p.curToken.Type != token.IDENT {
		return stmt, fmt.Errorf("Expected identifier after let, got %v", p.curToken)
	}

	stmt.Ident = ast.Identifier{
		Token: p.curToken,
		Name:  p.curToken.Literal,
	}

	p.nextToken()

	// assign (=)
	if p.curToken.Type != token.ASSIGN {
		return stmt, fmt.Errorf("Expected assignment (=), got %v", p.curToken)
	}

	p.nextToken()

	// expression
	expr, err := p.parseExpr(precedence.LOWEST)
	if err != nil {
		return stmt, err
	}

	stmt.Expr = expr

	// newline
	if !isEndOfStmt(p.curToken.Type) {
		return stmt, fmt.Errorf("Expected newline (\\n), got %v", p.curToken)
	}

	p.nextToken()

	return stmt, nil
}

func isEndOfStmt(tok token.TokenType) bool {
	return tok == token.NEWLINE || tok == token.EOF
}
