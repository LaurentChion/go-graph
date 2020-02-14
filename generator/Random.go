package generator

import (
	"github.com/LaurentChion/go-graph/core"
)

func RandomGraphGenerator(numberOfNode int, numberOfEdges int) *core.Graph {
	g := core.Graph{}

	for range make([]int, numberOfNode) {
		a := core.NewAttributes()
		g.AddNode(a) // no default attributes
	}

	for range make([]int, numberOfEdges) {
		a := core.NewAttributes()
		n1 := g.GetRandomNode()
		n2 := g.GetRandomNode()
		g.AddEdge(n1, n2, a) // no default attributes
	}
	return &g
}
