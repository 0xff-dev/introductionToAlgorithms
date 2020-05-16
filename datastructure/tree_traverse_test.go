package datastructure

import (
	"fmt"
	"testing"
)

func TestTreeNodeTraverse(t *testing.T) {
	tree := &treeNode{
		Data:  10,
		Left:  &treeNode{
			Data:  8,
			Left:  nil,
			Right: &treeNode{
				Data:  6,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &treeNode{
			Data:  12,
			Left:  &treeNode{
				Data:  5,
				Left:  nil,
				Right: nil,
			},
			Right: &treeNode{
				Data:  1,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println("pre order recursive -----")
	tree.PreOrder()
	fmt.Println("pre order --------")
	tree.PreOrder1()

	fmt.Println("in order recursive -----")
	tree.InOrder()
	fmt.Println("in order ------")
	tree.InOrder1()

	fmt.Println("post order recursive----- ")
	tree.PostOrder()
	fmt.Println("post order ------")
	tree.PostOrder1()
}