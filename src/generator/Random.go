package generator

import (
	"fmt"

	"github.com/LaurentChion/go-graph/src/implementations/list"
	"github.com/LaurentChion/go-graph/src/interfaces"
)

func RandomGraphGenerator(numberOfNode int) interfaces.Graph {
	g := list.GraphList
	for i := range make([]int, numberOfNode) {
		n := list.Node.New(uint64(i))
		g.AddNode(n)
	}
	fmt.Println(g)
	return g
}
