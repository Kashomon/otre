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
	if len(tree.Root().Children) != 0 {
		t.Errorf("Shouldn't any children initially")
	}
}

func TestAddNode(t *testing.T) {
	tree := NewMovetree()
	node := tree.newNode()
	if node.Variation() != 0 {
		t.Errorf("Should be the first variation")
	}
	if node.Depth() != 1 {
		t.Errorf("The Depth should be 1")
	}
	if len(tree.Root().Children) != 1 {
		t.Errorf("Should have 1 child at the root")
	}
	parent, err := node.Parent()
	if err != nil {
		t.Errorf("Couldn't retrieve the parent")
	}
	if parent != tree.Root() {
		t.Errorf("The new node's parent node must be set correctly")
	}
}
