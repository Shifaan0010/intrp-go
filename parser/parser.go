package parser

import (
	"fmt"
	"log/slog"
	"monkey-interpreter/ast"
	"monkey-interpreter/lexer"
	"monkey-interpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) (*Parser, error) {
	p := &Parser{
		l: l,
	}

	var err error

	p.curToken, err = l.NextToken()

	if err != nil {
		return p, err
	}

	p.peekToken, err = l.NextToken()

	return p, err
}

func (p *Parser) ParseProgram() (*ast.Program, error) {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}

	for p.curToken.Type != token.EOF {
		stmt, err := p.parseStatement()
		if err != nil {
			slog.Debug("failed to parse statement", "stmt", stmt, "err", err)

			return program, err
		}

		slog.Debug("parsed statement", "stmt", stmt)

		program.Statements = append(program.Statements, stmt)
	}

	return program, nil
}

func (p *Parser) ParseStatement() (ast.Statement, error) {
	return p.parseStatement()
}

func (p *Parser) IsAtEof() bool {
	return p.curToken.Type == token.EOF
}

func (p *Parser) nextToken() error {
	var err error = nil

	p.curToken = p.peekToken
	p.peekToken, err = p.l.NextToken()

	// fmt.Println("curToken:", p.curToken, "peekToken:", p.peekToken)

	return err
}

func (p *Parser) nextExpect(t token.TokenType) error {
	if p.peekToken.Type != t {
		return fmt.Errorf("Expected next token to be, got %v", p.curToken)
	}

	p.nextToken()

	return nil
}
