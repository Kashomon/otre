package core

type Movetree struct {
	root        *Node
	CurrentNode *Node
}

func InitMovetree() *Movetree {
	var newNode *Node = &Node{}
	return &Movetree{newNode, newNode}
}

// Retrieve the root node. This exists to prevent users from changing the root
// node.
func (m *Movetree) Root() *Node {
	return m.root
}

// Add a new child node to the CurrentNode.
func (m *Movetree) newNode() *Node {
	return m.CurrentNode.newChild()
}
