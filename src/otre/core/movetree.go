package core

type Movetree struct {
	root        *Node
	currentNode *Node
}

func NewMovetree() *Movetree {
	var newNode *Node = &Node{}
	return &Movetree{newNode, newNode}
}

// Retrieve the root node. This exists to prevent users from changing the root
// node.
func (m *Movetree) Root() *Node {
	return m.root
}

func (m *Movetree) CurrentNode() *Node {
	return m.currentNode
}

func (m *Movetree) FromRoot() *Movetree {
	return &Movetree{m.root, m.root}
}

// Add a new child node to the CurrentNode.
func (m *Movetree) NewNode() *Node {
	return m.CurrentNode().NewChild()
}

func (m *Movetree) MoveUp() *Movetree {
	parent, _ := m.CurrentNode().Parent()
	return &Movetree{m.root, parent}
}
