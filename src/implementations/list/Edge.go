package list

type Edge struct {
	from, to *Node
	Weight   float64
}

func (e *Edge) From() *Node {
	return e.from
}

func (e *Edge) to() *Node {
	return e.to
}
