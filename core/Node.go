package core

type Node struct {
	Inputs, Outputs []*Edge
	Attributes      []Attribute
}

func (n *Node) AddAttribute(a Attribute) {
	n.Attributes = append(n.Attributes, a)
}
