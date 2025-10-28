package lexer

import (
	"bufio"
	"io"
	"log/slog"

	"intrp-go/token"
)

type Lexer struct {
	input bufio.Reader
	buf   []byte

	isPeek bool
	peek   byte
}

func New(input bufio.Reader) *Lexer {
	return &Lexer{
		input: input,
		buf:   []byte{},
	}
}

func (l *Lexer) NextToken() (token.Token, error) {
	tok, err := l.nextToken()

	slog.Debug("read token", "tok", tok)

	return tok, err
}

func (l *Lexer) nextToken() (token.Token, error) {
	l.buf = l.buf[:0]
	// fmt.Println(l.buf)

	err := l.skipWhitespace()
	if err != nil {
		if err == io.EOF {
			return token.Token{Type: token.EOF, Literal: string([]byte{0})}, nil
		}
		return token.Token{Type: token.ILLEGAL}, err
	}

	err = l.readByte(false)
	if err != nil {
		if err == io.EOF {
			return token.Token{Type: token.EOF, Literal: string(l.buf)}, nil
		}
		return token.Token{Type: token.ILLEGAL}, err
	}

	b := l.buf[0]

	switch b {
	case '=':
		b, err := l.peekByte()
		if err == nil && b == '=' {
			l.readByte(false)
			return token.Token{Type: token.EQ, Literal: string(l.buf)}, nil
		}

		return token.Token{Type: token.ASSIGN, Literal: "="}, nil
	case '+':
		return token.Token{Type: token.PLUS, Literal: string(l.buf)}, nil
	case '-':
		return token.Token{Type: token.MINUS, Literal: string(l.buf)}, nil
	case '!':
		if b, err := l.peekByte(); err == nil && b == '=' {
			l.readByte(false)
			return token.Token{Type: token.NOT_EQ, Literal: string(l.buf)}, nil
		}
		return token.Token{Type: token.BANG, Literal: string(l.buf)}, nil

	case '*':
		return token.Token{Type: token.ASTERISK, Literal: string(l.buf)}, nil
	case '/':
		return token.Token{Type: token.SLASH, Literal: string(l.buf)}, nil
	case '<':
		return token.Token{Type: token.LT, Literal: string(l.buf)}, nil
	case '>':
		return token.Token{Type: token.GT, Literal: string(l.buf)}, nil
	case '\n':
		return token.Token{Type: token.NEWLINE, Literal: string(l.buf)}, nil
	case '(':
		return token.Token{Type: token.LPAREN, Literal: string(l.buf)}, nil
	case ')':
		return token.Token{Type: token.RPAREN, Literal: string(l.buf)}, nil
	case ',':
		return token.Token{Type: token.COMMA, Literal: string(l.buf)}, nil
	case '{':
		return token.Token{Type: token.LBRACE, Literal: string(l.buf)}, nil
	case '}':
		return token.Token{Type: token.RBRACE, Literal: string(l.buf)}, nil
	case 0:
		return token.Token{Type: token.EOF, Literal: string(l.buf)}, nil

	default:
		if isLetter(b) || isNumber(b) {
			l.readWord()

			if isLetter(b) {
				return token.Token{Type: token.KeywordType(string(l.buf)), Literal: string(l.buf)}, nil
			} else if isNumber(b) {
				return token.Token{Type: token.INT, Literal: string(l.buf)}, nil
			}
		}

		return token.Token{Type: token.ILLEGAL, Literal: string(l.buf)}, nil
	}
}

func (l *Lexer) peekByte() (byte, error) {
	if !l.isPeek {
		b, err := l.input.ReadByte()
		if err != nil {
			return b, err
		}

		l.peek = b
		l.isPeek = true
	}

	return l.peek, nil
}

func (l *Lexer) readByte(skip bool) error {
	var rdByte byte

	if l.isPeek {
		rdByte = l.peek
		l.isPeek = false
	} else {
		b, err := l.input.ReadByte()
		if err != nil {
			return err
		}

		rdByte = b
	}

	if !skip {
		l.buf = append(l.buf, rdByte)
	}

	return nil
}

func (l *Lexer) readWord() error {
	for {
		b, err := l.peekByte()
		if err != nil {
			return err
		}

		if isWhitespace(b) || !(isLetter(b) || isNumber(b)){
			break
		}

		err = l.readByte(false)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *Lexer) skipWhitespace() error {
	for {
		b, err := l.peekByte()
		if err != nil {
			return err
		}

		if !isWhitespace(b) {
			break
		}

		err = l.readByte(true)
		if err != nil {
			return err
		}

		// fmt.Printf("skip %q\n", b)
	}

	return nil
}

func isWhitespace(b byte) bool {
	return b == ' ' || b == '\r' || b == '\t'
}

func isLetter(b byte) bool {
	return (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z')
}

func isNumber(b byte) bool {
	return (b >= '0' && b <= '9')
}
