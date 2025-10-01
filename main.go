package main

import (
	"bufio"
	"fmt"
	"os"

	"monkey-interpreter/lexer"
	"monkey-interpreter/token"
)

func main() {
	l := lexer.New(*bufio.NewReader(os.Stdin))

	for {
		tok, err := l.NextToken()

		if err != nil {
			fmt.Println("err", err)
			break
		}

		fmt.Printf("%v\n", tok)
		// fmt.Println(l)

		if tok.Type == token.EOF {
			break
		}
	}
}
