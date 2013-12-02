package core

import "errors"

type Node struct {
	props     *Props
	variation int
	depth     int
	parent    *Node
	children  []*Node
}

// Get the parent node.
func (n *Node) Parent() (*Node, error) {
	if n.depth == 0 {
		return n, errors.New("At root")
	}
	return n.parent, nil
}

func (n *Node) Variation() int    { return n.variation }
func (n *Node) Depth() int        { return n.depth }
func (n *Node) Children() []*Node { return n.children }
func (n *Node) Props() *Props     { return n.props }

// Add a new Child node to this node. This should always succeed.
func (n *Node) NewChild() *Node {
	newChild := &Node{
		props:     NewProps(),
		variation: len(n.Children()),
		depth:     n.depth + 1,
		parent:    n,
		children:  make([]*Node, 0, 1),
	}
	n.children = append(n.children, newChild)
	return newChild
}

func NewNode() *Node {
	return &Node{
		NewProps(),
		0,
		0,
		nil,
		make([]*Node, 0, 1),
	}
}
