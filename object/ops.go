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

func Neg(a Object) (Object, error) {
	switch t := a.Type(); t {
	case INTEGER:
		aVal, _ := a.(*Integer)

		return &Integer{Val: -aVal.Val}, nil

	default:
		return nil, fmt.Errorf("Negate not implemented for type %s", t)
	}
}

func Not(a Object) (Object, error) {
	switch t := a.Type(); t {
	case BOOLEAN:
		aVal, _ := a.(*Boolean)

		return &Boolean{Val: !aVal.Val}, nil

	default:
		return nil, fmt.Errorf("Not not implemented for type %s", t)
	}
}

func Sub(a, b Object) (Object, error) {
	negB, err := Neg(b)
	if err != nil {
		return nil, err
	}

	return Add(a, negB)
}

func Mult(a, b Object) (Object, error) {
	if a.Type() != b.Type() {
		return nil, fmt.Errorf("expected right expr type to be %s, got %s", a.Type(), b.Type())
	}

	switch t := a.Type(); t {
	case INTEGER:
		aVal, _ := a.(*Integer)
		bVal, _ := b.(*Integer)

		return &Integer{Val: aVal.Val * bVal.Val}, nil

	default:
		return nil, fmt.Errorf("Mult not implemented for type %s", t)
	}
}

func Div(a, b Object) (Object, error) {
	if a.Type() != b.Type() {
		return nil, fmt.Errorf("expected right expr type to be %s, got %s", a.Type(), b.Type())
	}

	switch t := a.Type(); t {
	case INTEGER:
		aVal, _ := a.(*Integer)
		bVal, _ := b.(*Integer)

		return &Integer{Val: aVal.Val / bVal.Val}, nil

	default:
		return nil, fmt.Errorf("Div not implemented for type %s", t)
	}
}
