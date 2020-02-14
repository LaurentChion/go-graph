package generator

import (
	"fmt"

	"github.com/LaurentChion/go-graph/core"
)

func RandomGraphGenerator(numberOfNode int, numberOfEdges int) core.Graph {
	g := core.Graph{}

	for i := range make([]int, numberOfNode) {
		g.AddNode([]core.Attribute{
			core.Attribute{"id", fmt.Sprintf("%d", i)},
		})
	}

	for range make([]int, numberOfEdges) {
		n1 := g.GetRandomNode()
		n2 := g.GetRandomNode()
		g.AddEdge(n1, n2, nil)
	}
	return g
}
