package interfaces

type Edge interface {
	func (n *Edge) GetId() uint64
	func (n *Edge) GetEdges() []*Edges
}