package precedence

import (
	"intrp-go/token"
)

type Precedence int

const (
	_ = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // < or >
	COMMA
	SUM         // +
	PRODUCT     // *
	PREFIX      // +x or !x
	CALL        // foo()
)

func TokenPrecedence(tok token.TokenType) Precedence {
	switch tok {

	case token.COMMA:
		return COMMA

	case token.EQ:
		fallthrough
	case token.NOT_EQ:
		return EQUALS

	case token.LT:
		fallthrough
	case token.GT:
		return LESSGREATER

	case token.PLUS:
		fallthrough
	case token.MINUS:
		return SUM

	case token.ASTERISK:
		fallthrough
	case token.SLASH:
		return PRODUCT

	case token.LPAREN:
		return CALL
	}

	return LOWEST
}
