package object

import "fmt"

type Integer struct {
	Val int64
}

func (o *Integer) Type() Type {
	return INTEGER
}

func (o *Integer) Inspect() string {
	return o.String()
}

func (o *Integer) String() string {
	return fmt.Sprintf("%d", o.Val)
}
