package main

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"

	"intrp-go/ast"
	"intrp-go/eval"
	"intrp-go/lexer"
	"intrp-go/parser"
)

func readStmt(bufRead bufio.Reader) io.Reader {
	sb := strings.Builder{}

	fmt.Print("> ")

	for {
		line, _ := bufRead.ReadString('\n')
		
		if line == "\n" {
			break
		}

		sb.WriteString(line)

		fmt.Print("..")
	}

	return strings.NewReader(sb.String())
}

func main() {
	// opts := &slog.HandlerOptions{
	// 	Level: slog.LevelError,
	// }
	//
	// handler := slog.NewTextHandler(os.Stdout, opts)
	// logger := slog.New(handler)
	// slog.SetDefault(logger)

	bufRead := bufio.NewReader(os.Stdin)

	for {
		stmtRead := readStmt(*bufRead)

		l := lexer.New(*bufio.NewReader(stmtRead))

		p, err := parser.New(l)
		if err != nil {
			slog.Error("failed to init parser", "err", err)
			continue
		}

		stmt, err := p.ParseStatement()
		if err != nil {
			slog.Error("failed to parse statement", "err", err)
			continue
		}

		if _, ok := stmt.(*ast.EmptyStatement); ok {
			continue
		}

		// fmt.Printf("parsed statement: %#v\n", stmt)
		// fmt.Printf("parsed statement: %s\n", stmt)

		evald, err := eval.Eval(stmt)
		if err != nil {
			fmt.Println("error:", err)
		}

		fmt.Println(evald)
	}
}
