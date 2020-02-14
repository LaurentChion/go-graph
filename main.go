package main

import (
	"math/rand"

	"github.com/LaurentChion/go-graph/core"
	"github.com/LaurentChion/go-graph/generator"
)

func putSomeIdOnNode(g *core.Graph) {
	for i, n := range g.Nodes {
		n.Attributes.PutInt("id", i)
	}
}

func putSomeRandomWeightOnEdge(g *core.Graph) {

	for _, e := range g.Edges {
		e.Attributes.PutFloat64("weight", rand.Float64())
	}
}

func main() {
	// generate a graph
	g := generator.RandomGraphGenerator(10, 10)

	// add some label on it
	putSomeIdOnNode(g)
	putSomeRandomWeightOnEdge(g)

	// show the network
	g.Display()
}
