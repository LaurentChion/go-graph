package core

type Node struct {
	Inputs, Outputs []*Edge
	Attributes      *Attributes
}

// return the index of a node if it is in the collection
func (n1 *Node) isInputOf(n2 *Node) int {
	for i, e := range n1.Outputs {
		if e.To == n2 {
			return i
		}
	}
	return -1
}
func (n1 *Node) isOutputOf(n2 *Node) int {
	return n2.isInputOf(n1)
}
