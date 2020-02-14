package core

type Node struct {
	Inputs, Outputs []*Edge
	Attributes      *Attributes
}
