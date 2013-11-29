//
// Package core provides basic functionality for dealing with Go files and
// trees. Thus, this package contains a parser for SGFs and a tree of moves (a
// Movetree).  Movetrees are the basic data structure used to represent Go
// games and problems
//
// Creating 
//
// To create a new Movetree from an SGF:
// 	mt, err := FromSgfCoord("(;GM[1]...").parse()
// 
package core

type Movetree struct {
	root        *Node
	currentNode *Node
}

func NewMovetree() *Movetree {
	var newNode *Node = NewNode()
	return &Movetree{newNode, newNode}
}

// Retrieve the root node. This exists to prevent users from changing the root
// node.
func (m *Movetree) Root() *Node { return m.root }

// Retrieve the current node.
func (m *Movetree) Node() *Node { return m.currentNode }

// Return a new movetree that starts from the root node.
func (m *Movetree) FromRoot() *Movetree { return &Movetree{m.root, m.root} }

// Add a new child node to the CurrentNode.
func (m *Movetree) NewNode() *Node {
	child := m.Node().NewChild()
	m.currentNode = child
	return child
}

func (m *Movetree) MoveUp() *Movetree {
	parent, _ := m.Node().Parent()
	m.currentNode = parent
	return m
}

func (m *Movetree) MoveDown(v int) *Movetree {
	if len(m.Node().Children()) == 0 {
		return m
	}
	if v >= len(m.Node().Children()) {
		v = len(m.Node().Children()) - 1
	}
	m.currentNode = m.Node().Children()[v]
	return m
}
