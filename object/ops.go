package object

import "fmt"

func Add(a, b Object) (Object, error) {
	if a.Type() != b.Type() {
		return nil, fmt.Errorf("expected right expr type to be %s, got %s", a.Type(), b.Type())
	}

	switch t := a.Type(); t {
	case INTEGER:
		aVal, _ := a.(*Integer)
		bVal, _ := b.(*Integer)

		return &Integer{Val: aVal.Val + bVal.Val}, nil

	default:
		return nil, fmt.Errorf("Add not implemented for type %s", t)
	}
}
