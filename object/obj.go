package object

type Type int

const (
	UNKNOWN = iota
	NULL
	BOOLEAN
	INTEGER
	FN
)

func (t Type) String() string {
	switch t {
	case UNKNOWN:
		return "UNKNOWN"
	case NULL:
		return "NULL"
	case BOOLEAN:
		return "BOOLEAN"
	case INTEGER:
		return "INTEGER"
	case FN:
		return "FUNCTION"
	default:
		return "< >"
	}
}

type Object interface {
	Type() Type
	Inspect() string
}
