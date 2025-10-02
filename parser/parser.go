package parser

import (
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

func (p *Parser) ParseProgram() (ast.Program, error) {
	program := ast.Program{
		Statements: []ast.Statement{},
	}

	for p.curToken.Type != token.EOF {
		stmt, err := p.parseStatement()
		if err != nil {
			return program, err
		}

		slog.Debug("parsed statement", "stmt", stmt)

		program.Statements = append(program.Statements, stmt)
	}

	return program, nil
}

func (p *Parser) nextToken() error {
	var err error = nil

	p.curToken = p.peekToken
	p.peekToken, err = p.l.NextToken()

	return err
}
