package core

import (
	"fmt"
	"math/rand"
)

type Graph struct {
	Nodes []*Node
	Edges []*Edge
}

// Add a node
func (g *Graph) AddNode(attributes []Attribute) *Graph {
	n := Node{nil, nil, attributes}
	g.Nodes = append(g.Nodes, &n)
	return g
}

// Add an edge from n1 to n2
func (g *Graph) AddEdge(n1, n2 *Node, attributes []Attribute) *Graph {
	e := Edge{n1, n2, attributes}
	n1.Outputs = append(n1.Outputs, &e)
	n2.Inputs = append(n2.Inputs, &e)
	g.Edges = append(g.Edges, &e)
	return g
}

// Get a random node
func (g *Graph) GetRandomNode() *Node {
	return g.Nodes[rand.Intn(len(g.Nodes))]
}

// Get a random edge
func (g *Graph) GetRandomEdge() *Edge {
	return g.Edges[rand.Intn(len(g.Edges))]
}

func (g *Graph) Display() {
	fmt.Println("Nodes: (Input, Output, Attributes")
	for _, v := range g.Nodes {
		fmt.Println(v)
	}
	fmt.Println("Edges: (From, To, Attributes)")
	for _, v := range g.Edges {
		fmt.Println(v)
	}
}