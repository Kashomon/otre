// Package core provides basic functionality for dealing with Go files and
// trees. This package contains a parser for SGFs and a struct describing a tree
// of moves (a Movetree or a.k.a Go AST).  Movetrees are the basic data
// structure used to represent Baduk games and problems.
//
// Creating
//
// To create a new Movetree from an SGF:
// 	mt, err := FromSgfCoord("(;GM[1]...").parse()
//
package core

import (
	"strconv"
)

type Movetree struct {
	root        *Node
	currentNode *Node
}

// Create a new Movetree.
func NewMovetree() *Movetree {
	newNode := NewNode()
	mt := &Movetree{newNode, newNode}
	setMovetreeDefaults(mt)
	return mt
}

func setMovetreeDefaults(mt *Movetree) *Movetree {
	p := mt.Root().Props()
	if _, ok := p.Get(GM); !ok {
		p.Set(GM, "1")
	}
	if _, ok := p.Get(FF); !ok {
		p.Set(FF, "4")
	}
	if _, ok := p.Get(CA); !ok {
		p.Set(CA, "UTF-8")
	}
	if _, ok := p.Get(AP); !ok {
		p.Set(AP, "Otre")
	}
	if _, ok := p.Get(SZ); !ok {
		p.Set(SZ, "19")
	}
	if _, ok := p.Get(KM); !ok {
		p.Set(KM, "0.00")
	}
	if _, ok := p.Get(ST); !ok {
		p.Set(ST, "2")
	}
	return mt
}

// Create an empty Movetree.  This is useful for parsing, but most users will
// want to use the NewMovetree() version, which sets appropriate defaults on the
// root node.
func EmptyMovetree() *Movetree {
	var newNode *Node = NewNode()
	return &Movetree{newNode, newNode}
}

func (m *Movetree) Root() *Node   { return m.root }
func (m *Movetree) Node() *Node   { return m.currentNode }
func (m *Movetree) Props() *Props { return m.Node().Props() }

// Return a new movetree that starts from the root node.
func (m *Movetree) FromRoot() *Movetree { return &Movetree{m.root, m.root} }

// Add a new child node to the CurrentNode.
func (m *Movetree) NewNode() *Node {
	child := m.Node().NewChild()
	m.currentNode = child
	return child
}

// Move up. In other words, the current node via Node is replaced by the node's
// parent.
func (m *Movetree) MoveUp() *Movetree {
	parent, _ := m.Node().Parent()
	m.currentNode = parent
	return m
}

// Move down.  In other words, choose the v'th child (where v is the variation
// number 0 indexed, and move down.
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

// Get the number of intersections on the Baduk Board.
func (m *Movetree) Intersections() int {
	prop, ok := m.Node().Props().GetFirst(SZ)
	if !ok {
		return 19
	} else {
		out, err := strconv.Atoi(prop)
		if err != nil {
			return 19
		} else {
			return out
		}
	}
}
