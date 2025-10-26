package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"

	"monkey-interpreter/ast"
	"monkey-interpreter/lexer"
	"monkey-interpreter/parser"
)

func main() {
	// opts := &slog.HandlerOptions{
	//        Level: slog.LevelError,
	//    }
	//    handler := slog.NewTextHandler(os.Stdout, opts)
	//    logger := slog.New(handler)
	// slog.SetDefault(logger)

	fmt.Printf("> ")

	l := lexer.New(*bufio.NewReader(os.Stdin))

	p, err := parser.New(l)
	if err != nil {
		slog.Error("failed to init parser", "err", err)
	}

	for !p.IsAtEof() {
		stmt, err := p.ParseStatement()
		if err != nil {
			fmt.Println("failed to parse statement", "error:", err)
			break
		}

		if _, ok := stmt.(*ast.EmptyStatement); ok {
			continue
		}

		fmt.Printf("parsed statement: %#v\n", stmt)
		fmt.Printf("parsed statement: %s\n", stmt)

		fmt.Printf("> ")
	}
}
