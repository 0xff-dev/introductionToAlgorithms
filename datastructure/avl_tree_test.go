package datastructure

import "testing"

func TestNewAvlTree(t *testing.T) {
	tree := NewAvlTree()
	//for index := 1; index < 4; index ++ {
	//	tree.Insert(index)
	//}
	//tree.Display()
	for _, key := range []int{10, 6, 12, 9} {
		tree.Insert(key)
	}
	tree.Display()
	tree.DeleteNode(12)
	tree.Display()
}
