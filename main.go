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
	opts := &slog.HandlerOptions{
        Level: slog.LevelError,
    }
    handler := slog.NewTextHandler(os.Stdout, opts)
    logger := slog.New(handler)
	slog.SetDefault(logger)

	l := lexer.New(*bufio.NewReader(os.Stdin))

	p, err := parser.New(l)
	if err != nil {
		slog.Error("failed to init parser", "err", err)
	}

	program, err := p.ParseProgram()
	if err != nil {
		slog.Error("failed to parse program", "err:", err)
	}
	
	for _, stmt := range program.Statements {
		fmt.Printf("%#v\n", stmt)
		fmt.Printf("%s\n", stmt)
	}
}
