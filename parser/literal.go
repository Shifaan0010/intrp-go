package parser

import (
	"fmt"
	"intrp-go/ast"
	"intrp-go/token"
	"strconv"
)

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

func (p *Parser) parseBool() (ast.Expression, error) {
	// expect to only be called when curToken is int
	if p.curToken.Type != token.TRUE && p.curToken.Type != token.FALSE {
		panic(fmt.Sprintf("parseBool called with invalid token %v", p.curToken))
	}

	node := ast.BoolLiteral{
		Token: p.curToken,
	}

	if p.curToken.Type == token.TRUE {
		node.Val = true
	} else {
		node.Val = false
	}

	p.nextToken()

	return &node, nil
}
func (p *Parser) parseIdent() (*ast.Identifier, error) {
	// expect to only be called when curToken is ident
	if p.curToken.Type != token.IDENT {
		panic(fmt.Sprintf("parseIdent called with invalid token %v", p.curToken))
	}

	node := ast.Identifier{
		Name:  p.curToken.Literal,
		Token: p.curToken,
	}

	p.nextToken()

	return &node, nil
}
