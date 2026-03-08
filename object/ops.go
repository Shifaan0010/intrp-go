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

func Eq(a, b Object) (Object, error) {
	if a.Type() != b.Type() {
		return nil, fmt.Errorf("expected right expr type to be %s, got %s", a.Type(), b.Type())
	}

	switch t := a.Type(); t {
	case INTEGER:
		aVal, _ := a.(*Integer)
		bVal, _ := b.(*Integer)

		return &Boolean{Val: aVal.Val == bVal.Val}, nil

	case BOOLEAN:
		aVal, _ := a.(*Boolean)
		bVal, _ := b.(*Boolean)

		return &Boolean{Val: aVal.Val == bVal.Val}, nil

	default:
		return nil, fmt.Errorf("Eq not implemented for type %s", t)
	}
}

func Neq(a, b Object) (Object, error) {
	obj, err := Eq(a, b)
	obj.(*Boolean).Val = !obj.(*Boolean).Val
	return obj, err
}

func Lt(a, b Object) (Object, error) {
	if a.Type() != b.Type() {
		return nil, fmt.Errorf("expected right expr type to be %s, got %s", a.Type(), b.Type())
	}

	switch t := a.Type(); t {
	case INTEGER:
		aVal, _ := a.(*Integer)
		bVal, _ := b.(*Integer)

		return &Boolean{Val: aVal.Val < bVal.Val}, nil

	default:
		return nil, fmt.Errorf("Eq not implemented for type %s", t)
	}
}

func Lte(a, b Object) (Object, error) {
	objEq, err1 := Eq(a, b)
	if err1 != nil {
		return objEq, err1
	}

	objLt, err2 := Lt(a, b)
	if err2 != nil {
		return objLt, err2
	}

	return &Boolean{Val: objEq.(*Boolean).Val || objLt.(*Boolean).Val}, nil
}

func Gte(a, b Object) (Object, error) {
	objLt, err2 := Lt(a, b)
	if err2 != nil {
		return objLt, err2
	}

	return &Boolean{Val: !objLt.(*Boolean).Val}, nil
}

func Gt(a, b Object) (Object, error) {
	objEq, err1 := Eq(a, b)
	if err1 != nil {
		return objEq, err1
	}

	objLt, err2 := Lt(a, b)
	if err2 != nil {
		return objLt, err2
	}

	return &Boolean{Val: !objEq.(*Boolean).Val && !objLt.(*Boolean).Val}, nil
}
