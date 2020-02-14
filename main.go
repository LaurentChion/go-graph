package main

import (
	"fmt"
	"math/rand"

	"github.com/LaurentChion/go-graph/core"
	"github.com/LaurentChion/go-graph/generator"
)

func putSomeRandomWeightOnEdge(g *core.Graph) *core.Graph {
	for _, e := range g.Edges {
		e.AddAttribute(
			core.Attribute{"weight", fmt.Sprintf("%f", rand.Float64())},
		)
	}
	return g
}

func main() {
	g := generator.RandomGraphGenerator(10, 10)
	putSomeRandomWeightOnEdge(&g)
	g.Display()
}
