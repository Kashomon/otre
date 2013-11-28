package core

import "errors"

type Node struct {
	properties map[SgfProperty][]string
	variation  int
	depth      int
	parent     *Node
	Children   []*Node
}

// Get the parent node.
func (n *Node) Parent() (*Node, error) {
	if n.depth == 0 {
		return nil, errors.New("At root")
	}
	return n.parent, nil
}

// Get this node's variation number.
func (n *Node) Variation() int {
	return n.variation
}

// Get this node's depth number.
func (n *Node) Depth() int {
	return n.depth
}

func (n *Node) SetProp(property SgfProperty, data []string) {
	n.properties[property] = data
}

func (n *Node) AddToProp(property SgfProperty, data ...string) {
	props, ok := n.properties[property]
	if !ok {
		props = make([]string, 0, 10)
	}
	for _, d := range data {
		props = append(props, d)
	}
	n.properties[property] = props
}

// Add a new Child node to this node. This should always succeed.
func (n *Node) NewChild() *Node {
	newChild := &Node{
		variation: len(n.Children),
		parent:    n,
		depth:     n.depth + 1,
	}
	n.Children = append(n.Children, newChild)
	return newChild
}
