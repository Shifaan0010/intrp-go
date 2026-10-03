package object

import (
	"fmt"
	"intrp-go/ast"
)

type Fn struct {
	FnDecl *ast.FnDeclExpr
}

func (o *Fn) Type() Type {
	return FN
}

func (o *Fn) Inspect() string {
	return o.String()
}

func (o *Fn) String() string {
	return fmt.Sprintf("%s", o.FnDecl)
}
