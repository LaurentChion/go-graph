package main

import (
	"fmt"
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
	// generate a graph of n node and n(n-1) edge
	n := 50
	g := generator.RandomGraphGenerator(n, n*(n-1))

	// add some label on it
	putSomeIdOnNode(g)
	putSomeRandomWeightOnEdge(g)

	// show the network
	fmt.Println("The network")
	fmt.Println(g)
}
