package object

type Type int

const (
	UNKNOWN = iota
	NULL
	BOOLEAN
	INTEGER
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
	default:
		return "< >"
	}
}

type Object interface {
	Type() Type
	Inspect() string
}
