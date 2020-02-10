package interfaces

type Graph interface {
	/** Constructor **/
	// Default constructor with empty nodes and edges
	func (g *Graph) New() Graph

	/** Getter **/
	// Get nodes
	func (g *Graph) Nodes() []Node
	func (g *Graph) AddNode(n Node) *Graph
	
	// Get edges
	func (g *Graph) Edges() []Edge
	
	// Add a new Edges
	func (g *Graph) AddEdge(n1, n2 *Node, e *Edge) *Graph

	func (g *Graph) GetRandomNode() *Node
}