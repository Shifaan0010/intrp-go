package main

import (
	"fmt"
	"intrp-go/eval"
	"syscall/js"
)

var count int = 0

var env eval.Environment

func Eval(this js.Value, args []js.Value) any {
	count += 1

	fmt.Printf("inp: %q\n", args)

	if len(args) == 0 {
		return js.ValueOf("")
	}

	return js.ValueOf(env.Eval(args[0].String()))

	// return js.ValueOf(fmt.Sprint(count))
}

func main() {
	fmt.Println("Hello, WebAssembly!")

	js.Global().Set("goReplEval", js.FuncOf(Eval))

	wait := make(chan struct{}, 0)
	<-wait
}
