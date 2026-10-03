package eval_test

import (
	"bufio"
	"intrp-go/eval"
	"intrp-go/lexer"
	"intrp-go/object"
	"intrp-go/parser"
	"strings"
	"testing"
)

func TestEvalFn(t *testing.T) {
	const program = `
foo = fn(n) {
	if (n < 1) {
		1
	} else {
		n * foo(n - 1)
	}
}

foo(5)
`
	l := lexer.New(*bufio.NewReader(strings.NewReader(program)))
	p, err := parser.New(l)

	if err != nil {
		t.Fatalf("failed to initialize parser, err: %s", err)
	}

	prog, err := p.ParseProgram()

	if err != nil {
		t.Fatalf("failed to parse program, err: %s", err)
	}

	env := eval.NewEnv()

	var val object.Object = nil
	for _, stmt := range prog.Statements {
		val, err = env.EvalNode(stmt)
	}

	valInt, ok := val.(*object.Integer)

	if !ok || valInt.Val != 120 {
		t.Errorf("expected %d, got %v", 120, val)
	}
}
