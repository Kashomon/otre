package core

import "errors"

type Node struct {
	properties map[SgfProperty][]string
	variation  int
	depth      int
	parent     *Node
	children   []*Node
}

// Get the parent node.
func (n *Node) Parent() (*Node, error) {
	if n.depth == 0 {
		return n, errors.New("At root")
	}
	return n.parent, nil
}

// Get this node's variation number.
func (n *Node) Variation() int { return n.variation }

// Get this node's depth number.
func (n *Node) Depth() int        { return n.depth }
func (n *Node) Children() []*Node { return n.children }

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

func (n *Node) GetProps(p SgfProperty) ([]string, bool) {
	props, ok := n.properties[p]
	return props, ok
}

// Add a new Child node to this node. This should always succeed.
func (n *Node) NewChild() *Node {
	newChild := &Node{
		properties: make(map[SgfProperty][]string),
		variation:  len(n.Children()),
		depth:      n.depth + 1,
		parent:     n,
		children:   make([]*Node, 0, 1),
	}
	n.children = append(n.children, newChild)
	return newChild
}

func NewNode() *Node {
	return &Node{
		make(map[SgfProperty][]string),
		0,
		0,
		nil,
		make([]*Node, 0, 1),
	}
}
