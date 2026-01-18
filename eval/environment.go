package eval

import "fmt"
import "intrp-go/object"

type Environment struct {
	globals map[string]object.Object
	stack []map[string]object.Object
}

func NewEnv() Environment {
	return Environment{
		globals: map[string]object.Object{},
		stack: []map[string]object.Object{},
	}
}

func (e *Environment) currFrame() map[string]object.Object {
	if len(e.stack) == 0 {
		// return nil, errors.New("No stack frames")
		return e.globals
	}

	return e.stack[len(e.stack) - 1]
}

func (e *Environment) SetVal(name string, val object.Object) error {
	e.currFrame()[name] = val

	return nil
}

func (e *Environment) GetVal(name string) (object.Object, error) {
	val, ok := e.currFrame()[name]

	if !ok {
		globalVal, globalOk := e.globals[name]

		if !globalOk {
			return nil, fmt.Errorf("variable %q not set in current frame or globals", name)
		}

		return globalVal, nil
	}

	return val, nil
}
