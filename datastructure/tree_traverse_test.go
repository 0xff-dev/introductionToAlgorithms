package datastructure

import (
	"fmt"
	"testing"
)

func TestTreeNodeTraverse(t *testing.T) {
	tree := &treeNode{
		Data: 10,
		Left: &treeNode{
			Data: 8,
			Left: nil,
			Right: &treeNode{
				Data:  6,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &treeNode{
			Data: 12,
			Left: &treeNode{
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

	fmt.Println("floor traverse ----")
	tree.Floor()
}

func TestFindPostOrder(t *testing.T) {
	pre, in := []int{1, 2, 4, 7, 3, 5, 8, 9, 6}, []int{4, 7, 2, 1, 8, 5, 9, 3, 6}
	FindPostOrder(pre, in)
}

func TestTreeNode_Predecessor(t *testing.T) {
	tree := &treeNode{
		Data: 12,
		Left: &treeNode{
			Data: 10,
			Left: &treeNode{
				Data: 8,
				Left: nil,
				Right: &treeNode{
					Data:  9,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &treeNode{
				Data:  11,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &treeNode{
			Data: 15,
			Left: &treeNode{
				Data:  14,
				Left:  nil,
				Right: nil,
			},
			Right: &treeNode{
				Data:  16,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println("14 没有左子树切是父节点的左子树 aim is : 12")
	node := tree.Predecessor(14)
	fmt.Println(node.Data)
	fmt.Println("10 有自己的左子树: 9")
	node = tree.Predecessor(10)
	fmt.Println(node.Data)
	fmt.Println("11 没有左子树，是父节点的右子树 : 10")
	node = tree.Predecessor(11)
	fmt.Println(node.Data)

	// 后继
	fmt.Println("11 没有右子树切是父节点的右子树 : 12")
	node = tree.Successor(11)
	fmt.Println(node.Data)
	fmt.Println("8有自己右子树9")
	node = tree.Successor(8)
	fmt.Println(node.Data)
	fmt.Println("14是父节点的左子树: 15")
	node = tree.Successor(14)
	fmt.Println(node.Data)
}

func TestTreeNode_InsertNode(t *testing.T) {
	tree := treeNode{Data: 10}
	tree.InsertNode(6)
	tree.InsertNode(12)
	tree.Floor()
}

