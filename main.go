package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"

	"intrp-go/eval"
	"intrp-go/lexer"
	"intrp-go/parser"
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
		// j, _ := json.Marshal(stmt)
		// fmt.Printf("%s\n", string(j))

		// fmt.Printf("stmt: %#v", stmt)

		fmt.Printf("stmt: %s", stmt)

		evald, err := eval.EvalNode(stmt)
		if err != nil {
			fmt.Println("error:", err)
		}

		fmt.Println("evaled val:", evald)

		fmt.Println()
	}
}
