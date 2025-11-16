package parser

import (
	"bufio"
	"encoding/json"
	"intrp-go/lexer"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	const program = `let x = 10 * (2 + -3)
x + (y - 15)
123 + 25 - 10 * 40 < 5 == 1
!true

if 1 == 1 {
	2 + 2
} else {
	let x = 10
	!(x > 5)
}

let foo = fn (x, y, abc) {
	if x > y {
		return 1
	} else if y < x {
		return abc
	} else if abc + 1 == x - y {
		return 0
	} else {
		abc = abc + x + y
		return -abc
	}
}

1 == 1, 2 * 2, 3 < 5 + 4
(1 == 1), 2 * 2, (3 < 5) + 4

!(-123 + 25) - (!(10 * 40) < 5) == 1 + foo
foo(1, 2, 3)
foo(1, 2, 3 * foo(a, 1+1))
!(-123 + 25) - (!(10 * 40) < 5) == 1 + foo(1, 2, 3)`

	const expectedAstJson = `{"Statements":[{"Token":{"Type":21,"Literal":"let"},"Assign":{"Token":{"Type":0,"Literal":""},"Ident":{"Token":{"Type":2,"Literal":"x"},"Name":"x"},"Expr":{"Op":{"Type":8,"Literal":"*"},"Left":{"Token":{"Type":3,"Literal":"10"},"Val":10},"Right":{"Op":{"Type":5,"Literal":"+"},"Left":{"Token":{"Type":3,"Literal":"2"},"Val":2},"Right":{"Op":{"Type":6,"Literal":"-"},"Expr":{"Token":{"Type":3,"Literal":"3"},"Val":3}}}}}},{"Token":{"Type":0,"Literal":""},"Expr":{"Op":{"Type":5,"Literal":"+"},"Left":{"Token":{"Type":2,"Literal":"x"},"Name":"x"},"Right":{"Op":{"Type":6,"Literal":"-"},"Left":{"Token":{"Type":2,"Literal":"y"},"Name":"y"},"Right":{"Token":{"Type":3,"Literal":"15"},"Val":15}}}},{"Token":{"Type":0,"Literal":""},"Expr":{"Op":{"Type":12,"Literal":"=="},"Left":{"Op":{"Type":10,"Literal":"\u003c"},"Left":{"Op":{"Type":6,"Literal":"-"},"Left":{"Op":{"Type":5,"Literal":"+"},"Left":{"Token":{"Type":3,"Literal":"123"},"Val":123},"Right":{"Token":{"Type":3,"Literal":"25"},"Val":25}},"Right":{"Op":{"Type":8,"Literal":"*"},"Left":{"Token":{"Type":3,"Literal":"10"},"Val":10},"Right":{"Token":{"Type":3,"Literal":"40"},"Val":40}}},"Right":{"Token":{"Type":3,"Literal":"5"},"Val":5}},"Right":{"Token":{"Type":3,"Literal":"1"},"Val":1}}},{"Token":{"Type":0,"Literal":""},"Expr":{"Op":{"Type":7,"Literal":"!"},"Expr":{"Token":{"Type":22,"Literal":"true"},"Val":true}}},{"Token":{"Type":15,"Literal":"\n"}},{"Token":{"Type":0,"Literal":""},"Expr":{"Cond":{"Op":{"Type":12,"Literal":"=="},"Left":{"Token":{"Type":3,"Literal":"1"},"Val":1},"Right":{"Token":{"Type":3,"Literal":"1"},"Val":1}},"If":{"Stmts":[{"Token":{"Type":15,"Literal":"\n"}},{"Token":{"Type":0,"Literal":""},"Expr":{"Op":{"Type":5,"Literal":"+"},"Left":{"Token":{"Type":3,"Literal":"2"},"Val":2},"Right":{"Token":{"Type":3,"Literal":"2"},"Val":2}}}],"Tok":{"Type":0,"Literal":""}},"Else":{"Stmts":[{"Token":{"Type":15,"Literal":"\n"}},{"Token":{"Type":21,"Literal":"let"},"Assign":{"Token":{"Type":0,"Literal":""},"Ident":{"Token":{"Type":2,"Literal":"x"},"Name":"x"},"Expr":{"Token":{"Type":3,"Literal":"10"},"Val":10}}},{"Token":{"Type":0,"Literal":""},"Expr":{"Op":{"Type":7,"Literal":"!"},"Expr":{"Op":{"Type":11,"Literal":"\u003e"},"Left":{"Token":{"Type":2,"Literal":"x"},"Name":"x"},"Right":{"Token":{"Type":3,"Literal":"5"},"Val":5}}}}],"Tok":{"Type":0,"Literal":""}},"Tok":{"Type":24,"Literal":"if"}}},{"Token":{"Type":15,"Literal":"\n"}},{"Token":{"Type":21,"Literal":"let"},"Assign":{"Token":{"Type":0,"Literal":""},"Ident":{"Token":{"Type":2,"Literal":"foo"},"Name":"foo"},"Expr":{"Params":[{"Token":{"Type":2,"Literal":"x"},"Name":"x"},{"Token":{"Type":2,"Literal":"y"},"Name":"y"},{"Token":{"Type":2,"Literal":"abc"},"Name":"abc"}],"Block":{"Stmts":[{"Token":{"Type":15,"Literal":"\n"}},{"Token":{"Type":0,"Literal":""},"Expr":{"Cond":{"Op":{"Type":11,"Literal":"\u003e"},"Left":{"Token":{"Type":2,"Literal":"x"},"Name":"x"},"Right":{"Token":{"Type":2,"Literal":"y"},"Name":"y"}},"If":{"Stmts":[{"Token":{"Type":15,"Literal":"\n"}},{"Token":{"Type":26,"Literal":"return"},"Expr":{"Token":{"Type":3,"Literal":"1"},"Val":1}}],"Tok":{"Type":0,"Literal":""}},"Else":{"Cond":{"Op":{"Type":10,"Literal":"\u003c"},"Left":{"Token":{"Type":2,"Literal":"y"},"Name":"y"},"Right":{"Token":{"Type":2,"Literal":"x"},"Name":"x"}},"If":{"Stmts":[{"Token":{"Type":15,"Literal":"\n"}},{"Token":{"Type":26,"Literal":"return"},"Expr":{"Token":{"Type":2,"Literal":"abc"},"Name":"abc"}}],"Tok":{"Type":0,"Literal":""}},"Else":{"Cond":{"Op":{"Type":12,"Literal":"=="},"Left":{"Op":{"Type":5,"Literal":"+"},"Left":{"Token":{"Type":2,"Literal":"abc"},"Name":"abc"},"Right":{"Token":{"Type":3,"Literal":"1"},"Val":1}},"Right":{"Op":{"Type":6,"Literal":"-"},"Left":{"Token":{"Type":2,"Literal":"x"},"Name":"x"},"Right":{"Token":{"Type":2,"Literal":"y"},"Name":"y"}}},"If":{"Stmts":[{"Token":{"Type":15,"Literal":"\n"}},{"Token":{"Type":26,"Literal":"return"},"Expr":{"Token":{"Type":3,"Literal":"0"},"Val":0}}],"Tok":{"Type":0,"Literal":""}},"Else":{"Stmts":[{"Token":{"Type":15,"Literal":"\n"}},{"Token":{"Type":0,"Literal":""},"Ident":{"Token":{"Type":2,"Literal":"abc"},"Name":"abc"},"Expr":{"Op":{"Type":5,"Literal":"+"},"Left":{"Op":{"Type":5,"Literal":"+"},"Left":{"Token":{"Type":2,"Literal":"abc"},"Name":"abc"},"Right":{"Token":{"Type":2,"Literal":"x"},"Name":"x"}},"Right":{"Token":{"Type":2,"Literal":"y"},"Name":"y"}}},{"Token":{"Type":26,"Literal":"return"},"Expr":{"Op":{"Type":6,"Literal":"-"},"Expr":{"Token":{"Type":2,"Literal":"abc"},"Name":"abc"}}}],"Tok":{"Type":0,"Literal":""}},"Tok":{"Type":24,"Literal":"if"}},"Tok":{"Type":24,"Literal":"if"}},"Tok":{"Type":24,"Literal":"if"}}}],"Tok":{"Type":0,"Literal":""}},"Tok":{"Type":20,"Literal":"fn"}}}},{"Token":{"Type":15,"Literal":"\n"}},{"Token":{"Type":0,"Literal":""},"Expr":{"Op":{"Type":12,"Literal":"=="},"Left":{"Token":{"Type":3,"Literal":"1"},"Val":1},"Right":{"Op":{"Type":10,"Literal":"\u003c"},"Left":{"Op":{"Type":14,"Literal":","},"Left":{"Op":{"Type":14,"Literal":","},"Left":{"Token":{"Type":3,"Literal":"1"},"Val":1},"Right":{"Op":{"Type":8,"Literal":"*"},"Left":{"Token":{"Type":3,"Literal":"2"},"Val":2},"Right":{"Token":{"Type":3,"Literal":"2"},"Val":2}}},"Right":{"Token":{"Type":3,"Literal":"3"},"Val":3}},"Right":{"Op":{"Type":5,"Literal":"+"},"Left":{"Token":{"Type":3,"Literal":"5"},"Val":5},"Right":{"Token":{"Type":3,"Literal":"4"},"Val":4}}}}},{"Token":{"Type":0,"Literal":""},"Expr":{"Op":{"Type":14,"Literal":","},"Left":{"Op":{"Type":14,"Literal":","},"Left":{"Op":{"Type":12,"Literal":"=="},"Left":{"Token":{"Type":3,"Literal":"1"},"Val":1},"Right":{"Token":{"Type":3,"Literal":"1"},"Val":1}},"Right":{"Op":{"Type":8,"Literal":"*"},"Left":{"Token":{"Type":3,"Literal":"2"},"Val":2},"Right":{"Token":{"Type":3,"Literal":"2"},"Val":2}}},"Right":{"Op":{"Type":5,"Literal":"+"},"Left":{"Op":{"Type":10,"Literal":"\u003c"},"Left":{"Token":{"Type":3,"Literal":"3"},"Val":3},"Right":{"Token":{"Type":3,"Literal":"5"},"Val":5}},"Right":{"Token":{"Type":3,"Literal":"4"},"Val":4}}}},{"Token":{"Type":15,"Literal":"\n"}},{"Token":{"Type":0,"Literal":""},"Expr":{"Op":{"Type":12,"Literal":"=="},"Left":{"Op":{"Type":6,"Literal":"-"},"Left":{"Op":{"Type":7,"Literal":"!"},"Expr":{"Op":{"Type":5,"Literal":"+"},"Left":{"Op":{"Type":6,"Literal":"-"},"Expr":{"Token":{"Type":3,"Literal":"123"},"Val":123}},"Right":{"Token":{"Type":3,"Literal":"25"},"Val":25}}},"Right":{"Op":{"Type":10,"Literal":"\u003c"},"Left":{"Op":{"Type":7,"Literal":"!"},"Expr":{"Op":{"Type":8,"Literal":"*"},"Left":{"Token":{"Type":3,"Literal":"10"},"Val":10},"Right":{"Token":{"Type":3,"Literal":"40"},"Val":40}}},"Right":{"Token":{"Type":3,"Literal":"5"},"Val":5}}},"Right":{"Op":{"Type":5,"Literal":"+"},"Left":{"Token":{"Type":3,"Literal":"1"},"Val":1},"Right":{"Token":{"Type":2,"Literal":"foo"},"Name":"foo"}}}},{"Token":{"Type":0,"Literal":""},"Expr":{"Op":{"Type":16,"Literal":"("},"Left":{"Token":{"Type":2,"Literal":"foo"},"Name":"foo"},"Right":{"Op":{"Type":14,"Literal":","},"Left":{"Op":{"Type":14,"Literal":","},"Left":{"Token":{"Type":3,"Literal":"1"},"Val":1},"Right":{"Token":{"Type":3,"Literal":"2"},"Val":2}},"Right":{"Token":{"Type":3,"Literal":"3"},"Val":3}}}},{"Token":{"Type":0,"Literal":""},"Expr":{"Op":{"Type":16,"Literal":"("},"Left":{"Token":{"Type":2,"Literal":"foo"},"Name":"foo"},"Right":{"Op":{"Type":14,"Literal":","},"Left":{"Op":{"Type":14,"Literal":","},"Left":{"Token":{"Type":3,"Literal":"1"},"Val":1},"Right":{"Token":{"Type":3,"Literal":"2"},"Val":2}},"Right":{"Op":{"Type":8,"Literal":"*"},"Left":{"Token":{"Type":3,"Literal":"3"},"Val":3},"Right":{"Op":{"Type":16,"Literal":"("},"Left":{"Token":{"Type":2,"Literal":"foo"},"Name":"foo"},"Right":{"Op":{"Type":14,"Literal":","},"Left":{"Token":{"Type":2,"Literal":"a"},"Name":"a"},"Right":{"Op":{"Type":5,"Literal":"+"},"Left":{"Token":{"Type":3,"Literal":"1"},"Val":1},"Right":{"Token":{"Type":3,"Literal":"1"},"Val":1}}}}}}}},{"Token":{"Type":0,"Literal":""},"Expr":{"Op":{"Type":12,"Literal":"=="},"Left":{"Op":{"Type":6,"Literal":"-"},"Left":{"Op":{"Type":7,"Literal":"!"},"Expr":{"Op":{"Type":5,"Literal":"+"},"Left":{"Op":{"Type":6,"Literal":"-"},"Expr":{"Token":{"Type":3,"Literal":"123"},"Val":123}},"Right":{"Token":{"Type":3,"Literal":"25"},"Val":25}}},"Right":{"Op":{"Type":10,"Literal":"\u003c"},"Left":{"Op":{"Type":7,"Literal":"!"},"Expr":{"Op":{"Type":8,"Literal":"*"},"Left":{"Token":{"Type":3,"Literal":"10"},"Val":10},"Right":{"Token":{"Type":3,"Literal":"40"},"Val":40}}},"Right":{"Token":{"Type":3,"Literal":"5"},"Val":5}}},"Right":{"Op":{"Type":5,"Literal":"+"},"Left":{"Token":{"Type":3,"Literal":"1"},"Val":1},"Right":{"Op":{"Type":16,"Literal":"("},"Left":{"Token":{"Type":2,"Literal":"foo"},"Name":"foo"},"Right":{"Op":{"Type":14,"Literal":","},"Left":{"Op":{"Type":14,"Literal":","},"Left":{"Token":{"Type":3,"Literal":"1"},"Val":1},"Right":{"Token":{"Type":3,"Literal":"2"},"Val":2}},"Right":{"Token":{"Type":3,"Literal":"3"},"Val":3}}}}}}]}`

	// var expected_prog ast.Program
	// err := json.Unmarshal([]byte(ast_json), expected_prog)
	// if err != nil {
	// 	t.Fatalf("failed to unmarshal expected ast, err: %s", err)
	// }

	lexer := lexer.New(*bufio.NewReader(strings.NewReader(program)))
	parser, err := New(lexer)

	if err != nil {
		t.Fatalf("failed to initialize parser, err: %s", err)
	}

	prog, err := parser.ParseProgram()

	if err != nil {
		t.Fatalf("failed to parse program, err: %s", err)
	}
	
	astJson, err := json.Marshal(prog)
	if err != nil {
		t.Fatalf("failed to marshal ast, err: %s", err)
	}

	if string(astJson) != expectedAstJson {
		t.Fatalf("ast json != expected ast\n%s\n%s", astJson, expectedAstJson)
	}
}
