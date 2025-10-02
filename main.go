package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"

	"monkey-interpreter/lexer"
	"monkey-interpreter/parser"
)

func main() {
	l := lexer.New(*bufio.NewReader(os.Stdin))

	p, err := parser.New(l)
	if err != nil {
		slog.Error("failed to init parser", "err", err)
	}

	program, err := p.ParseProgram()
	if err != nil {
		slog.Error("failed to parse program", "err:", err)
	}
	
	fmt.Printf("program: %#v", program)
}
