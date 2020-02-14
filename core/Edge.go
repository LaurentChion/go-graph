package core

type Edge struct {
	From, To   *Node
	Attributes []Attribute
}

func (e *Edge) AddAttribute(a Attribute) {
	e.Attributes = append(e.Attributes, a)
}
