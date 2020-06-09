package datastructure

import (
	"fmt"
	"testing"
)

func TestNewRBTree(t *testing.T) {
	tree := NewRBTree()
	tree.Insert(10)
	tree.Insert(12)
	tree.Insert(11)
	tree.Insert(15)
	tree.Insert(2)
	tree.Root.Floor()

	self := tree.Root.Search(11)
	if self == nil {
		fmt.Println("11 not found")
	} else {
		tree.Delete(self)
		tree.Root.Floor()
	}
}
