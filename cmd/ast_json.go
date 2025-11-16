package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"intrp-go/lexer"
	"intrp-go/parser"
)

func main() {
	// opts := &slog.HandlerOptions{
	// 	Level: slog.LevelError,
	// }
	//
	// handler := slog.NewTextHandler(os.Stdout, opts)
	// logger := slog.New(handler)
	// slog.SetDefault(logger)

	l := lexer.New(*bufio.NewReader(os.Stdin))

	p, err := parser.New(l)
	if err != nil {
		slog.Error("failed to init parser", "err", err)
	}

	program, err := p.ParseProgram()
	if err != nil {
		slog.Error("failed to parse program", "err:", err)
	}

	ast_json, err := json.Marshal(program)
	if err != nil {
		slog.Error("failed to marshal program", "err:", err)
	}

	fmt.Println(string(ast_json))
}
