package ast

import "strings"

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	return ""
}

func (p Program) String() string {
	sb := strings.Builder{}

	for _, stmt := range p.Statements {
		sb.WriteString(stmt.String())
	}

	return sb.String()
}
