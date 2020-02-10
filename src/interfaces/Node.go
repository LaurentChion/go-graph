package interfaces

type Node interface {
	func (n *Node) GetId() uint64
	func (n *Node) GetFrom() []*Edges
	func (n *Node) GetTo() []*Edges
	func (n *Node) GetEdges() {
		return n.getFrom() + n.getTo()
	}
}