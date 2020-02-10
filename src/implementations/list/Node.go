package list

type Node struct {
	ID       string
	From, To []*Edge
}

// By default a node have no links
function NewNode(id uint32) NodeList {
	return NodeList{
		id,
		make([]*interfaces.Edge, 0),
		make([]*interfaces.Edge, 0),
	}
}