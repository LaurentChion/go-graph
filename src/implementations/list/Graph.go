package list

/* Graph interface should be abble to :
- add some node
- add some edges
This default implementation use a list
*/
type Graph struct {
	Nodes []graph.Node
	Edges []graph.Edge
}

// mutate both nodes
func (g *Graph) AddNode(n Node) *Graph {
	g.nodes = append(g.nodes, n)
}

// mutate both nodes
func (g *Graph) AddEdge(n1, n2 *Node, e *Edge) *Graph {
	n1.To = append(n1.To, e)
	n2.From = append(n2.From, e)
	g.Edges = append(g.Edges, e)
	return g
}
