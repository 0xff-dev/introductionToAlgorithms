package datastructure

import "testing"

func TestNewRBTree(t *testing.T) {
	tree := NewRBTree()
	tree.Insert(10)
	tree.Insert(12)
	tree.Insert(11)
	tree.Root.Floor()
}
