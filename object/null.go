package object

type Null struct{}

func (o *Null) Type() Type {
	return NULL
}

func (o *Null) Inspect() string {
	return o.String()
}

func (o *Null) String() string {
	return "null"
}
