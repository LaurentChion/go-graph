package core

import (
	"fmt"
	"math/rand"
)

type Graph struct {
	Nodes []*Node
	Edges []*Edge
}

func NewGraph() *Graph {
	return &Graph{}
}

// Add a node
func (g *Graph) AddNode(attributes *Attributes) *Graph {
	n := Node{nil, nil, attributes}
	g.Nodes = append(g.Nodes, &n)
	return g
}

// Add an edge from n1 to n2
func (g *Graph) AddEdge(n1, n2 *Node, attributes *Attributes) *Graph {
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

func (g Graph) String() string {

	numberOfNode := fmt.Sprintf("%d", len(g.Nodes)-1)
	numberOfLeadingZeros := len(numberOfNode)
	leadingZeroCode := fmt.Sprintf("0%d", numberOfLeadingZeros)
	s := fmt.Sprintf("%*s|", numberOfLeadingZeros+1, "")
	for i := range g.Nodes {
		s += fmt.Sprintf("%"+leadingZeroCode+"d|", i)
	}
	s += fmt.Sprintln()
	for i, n1 := range g.Nodes {
		s += fmt.Sprintf("%"+leadingZeroCode+"d ", i)
		for j, n2 := range g.Nodes {
			if j <= i {
				// if i is linked to j
				isOutput := n1.isOutputOf(n2) >= 0
				isInput := n1.isInputOf(n2) >= 0

				if isInput && isOutput && i == j {
					s += fmt.Sprintf("|🔃%*s", numberOfLeadingZeros-1, "")
				} else if isInput && isOutput {
					s += fmt.Sprintf("|↔%*s", numberOfLeadingZeros-1, "")
				} else if isInput {
					s += fmt.Sprintf("|⬆️%*s", numberOfLeadingZeros-1, "")
				} else if isOutput {
					s += fmt.Sprintf("|⬅️%*s", numberOfLeadingZeros-1, "")
				} else {
					s += fmt.Sprintf("|%*s", numberOfLeadingZeros, "")
				}
			}
		}
		s += fmt.Sprintln()
	}
	return s
}
