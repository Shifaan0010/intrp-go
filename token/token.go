package token

import "fmt"

type Token struct {
	Type    TokenType
	Literal string
}

func (t Token) String() string {
	return fmt.Sprintf("%q", t.Literal)
}

func KeywordType(s string) TokenType {
	switch s {
	case "fn":
		return FUNCTION
	case "let":
		return LET
	case "true":
		return TRUE
	case "false":
		return FALSE
	case "if":
		return IF
	case "else":
		return ELSE
	case "return":
		return RETURN
	default:
		return IDENT
	}
}

type TokenType int

const (
	ILLEGAL = iota
	EOF

	// Identifiers + literals
	IDENT
	INT

	// Operators
	ASSIGN   // =
	PLUS     // +
	MINUS    // -
	BANG     // !
	ASTERISK // *
	SLASH    // /

	LT // <
	GT // >

	EQ     // ==
	NOT_EQ // !=

	// Delimiters
	COMMA   // ,
	NEWLINE // \n
	LPAREN  // (
	RPAREN  // )
	LBRACE  // {
	RBRACE  // }

	// Keywords
	FUNCTION // fn
	LET      // let
	TRUE     // true
	FALSE    // false
	IF       // if
	ELSE     // else
	RETURN   // return
)

func (t TokenType) String() string {
	switch t {
	case ILLEGAL:
		return "ILLEGAL"
	case EOF:
		return "EOF"

	// Identifiers + literals
	case IDENT:
		return "IDENT"
	case INT:
		return "INT"

	// Operators
	case ASSIGN: // =
		return "ASSIGN"
	case PLUS: // +
		return "PLUS"
	case MINUS: // -
		return "MINUS"
	case BANG: // !
		return "BANG"
	case ASTERISK: // *
		return "ASTERISK"
	case SLASH: // /
		return "SLASH"

	case LT: // <
		return "LT"
	case GT: // >
		return "GT"

	case EQ: // ==
		return "EQ"
	case NOT_EQ: // !=
		return "NOT_EQ"

	// Delimiters
	case COMMA: // ,
		return "COMMA"
	case NEWLINE: // \n
		return "NEWLINE"
	case LPAREN: // (
		return "LPAREN"
	case RPAREN: // )
		return "RPAREN"
	case LBRACE: // {
		return "LBRACE"
	case RBRACE: // }
		return "RBRACE"

	// Keywords
	case FUNCTION: // fn
		return "FUNCTION"
	case LET: // let
		return "LET"
	case TRUE: // true
		return "TRUE"
	case FALSE: // false
		return "FALSE"
	case IF: // if
		return "IF"
	case ELSE: // else
		return "ELSE"
	case RETURN: // return
		return "RETURN"

	default:
		return "UNKNOWN"
	}
}
