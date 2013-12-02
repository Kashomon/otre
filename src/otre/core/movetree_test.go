package core

import "testing"

func TestBasicCreation(t *testing.T) {
	var tree *Movetree = NewMovetree()
	if tree == nil {
		t.Errorf("Shouldn't be nil")
	}
	if tree.Root().Variation() != 0 {
		t.Errorf("Variation should be initialized to 0")
	}
	if tree.Root().Depth() != 0 {
		t.Errorf("Depth should be initialized to 0")
	}
	if len(tree.Root().Children()) != 0 {
		t.Errorf("Shouldn't any children initially")
	}
	if gm, _ := tree.Props().GetFirst(GM); gm != "1" {
		t.Errorf("Expected 1, instead found %v", gm)
	}
	if ints := tree.Intersections(); ints != 19 {
		t.Errorf("Expected 19, instead found %v", ints)
	}
}

func TestAddNode(t *testing.T) {
	mt := NewMovetree()
	node := mt.NewNode()
	if node.Variation() != 0 {
		t.Errorf("Should be the first variation")
	}
	if node.Depth() != 1 {
		t.Errorf("The Depth should be 1")
	}
	if len(mt.Root().Children()) != 1 {
		t.Errorf("Should have 1 child at the root")
	}
	parent, err := node.Parent()
	if err != nil {
		t.Errorf("Couldn't retrieve the parent")
	}
	if parent != mt.Root() {
		t.Errorf("The new node's parent node must be set correctly")
	}
	if node != mt.Node() {
		t.Errorf("Current node not equal to child node")
	}
}

func TestMoveUpAndDown(t *testing.T) {
	mt := NewMovetree()
	mt.NewNode()
	mt.MoveUp()
	mt.NewNode()
	mt.NewNode()
	mt = mt.FromRoot()
	if c := len(mt.Node().Children()); c != 2 {
		t.Errorf("Expected two children, saw %v", c)
	}
	mt.MoveDown(1)
	if c := len(mt.Node().Children()); c != 1 {
		t.Errorf("Expected one child, saw %v", c)
	}
	mt.MoveDown(1)
	if c := len(mt.Node().Children()); c != 0 {
		t.Errorf("Expected no children, saw %v", c)
	}
	cnode := mt.Node()
	mt.MoveDown(0)
	if cnode != mt.MoveDown(0).Node() {
		t.Errorf("Expected to not move at the end.")
	}
	mt.MoveUp()
	if c := len(mt.Node().Children()); c != 1 {
		t.Errorf("Expected one child, saw %v", c)
	}
	mt.MoveUp()
	if c := len(mt.Node().Children()); c != 2 {
		t.Errorf("Expected two children, saw %v", c)
	}
	cnode = mt.Node()
	mt.MoveUp()
	if cnode != mt.MoveUp().Node() {
		t.Errorf("Expected to not move at the beginning.")
	}
}
