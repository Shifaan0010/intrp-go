package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"intrp-go/eval"
)

func readStmt(bufRead bufio.Reader) string {
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

	return sb.String()
}

func main() {
	// opts := &slog.HandlerOptions{
	// 	Level: slog.LevelError,
	// }
	//
	// handler := slog.NewTextHandler(os.Stdout, opts)
	// logger := slog.New(handler)
	// slog.SetDefault(logger)

	env := eval.NewEnv()

	bufRead := bufio.NewReader(os.Stdin)

	for {
		stmtStr := readStmt(*bufRead)

		fmt.Println(env.Eval(stmtStr))
	}
}
