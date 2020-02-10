package main

import (
	"fmt"

	"github.com/LaurentChion/go-graph/src/generator"
)

/*func putSomeRandomWeightedEdge(g *interfaces.Graph, nbOfEdges int) *interfaces.Graph {
	for range make([]int, nbOfEdges) {
		n1 := g.GetRandomNode()
		n2 := g.GetRandomNode()
		e := graph.Edge{
			n1, n2,
			rand.Float64(),
		}
		graph.AddEdge(n1, n2, &e)
	}
	return g
}*/

func main() {
	g := generator.RandomGraphGenerator(10)
	//putSomeRandomWeightedEdge(&g, 10)
	fmt.Println(g)
}
