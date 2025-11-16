package object

import "fmt"

type Boolean struct {
	Val bool
}

func (o *Boolean) Type() Type {
	return BOOLEAN
}

func (o *Boolean) Inspect() string {
	return o.String()
}

func (o *Boolean) String() string {
	return fmt.Sprintf("%t", o.Val)
}
