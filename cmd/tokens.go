package main

import (
	"bufio"
	"fmt"
	"os"

	"intrp-go/lexer"
	"intrp-go/token"
)

func main() {
	l := lexer.New(*bufio.NewReader(os.Stdin))

	fmt.Print("> ")
	for {
		tok, err := l.NextToken()

		if err != nil {
			fmt.Println("err", err)
			break
		}

		fmt.Printf("%v\n", tok.DbgString())
		// fmt.Println(l)

		if tok.Type == token.EOF {
			break
		} else if tok.Type == token.NEWLINE {
			fmt.Print("> ")
		}

	}
}
