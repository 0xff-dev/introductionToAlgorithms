package challengeprogrammingdatastructure

import "fmt"

type BSTNode struct {
	Val         int
	Left, Right *BSTNode
}

func BinarySearchTreeInsert(root *BSTNode, val int) *BSTNode {
	n := &BSTNode{Val: val}
	if root == nil {
		return n
	}

	if root.Val > val {
		root.Left = BinarySearchTreeInsert(root.Left, val)
	} else if root.Val < val {
		root.Right = BinarySearchTreeInsert(root.Right, val)
	}
	return root
}

func inOrder92(root *BSTNode) {
	if root.Left != nil {
		inOrder92(root.Left)
	}
	fmt.Print(root.Val, " ")
	if root.Right != nil {
		inOrder92(root.Right)
	}
}

func preOrder92(root *BSTNode) {
	fmt.Print(root.Val, " ")
	if root.Left != nil {
		preOrder92(root.Left)
	}
	if root.Right != nil {
		preOrder92(root.Right)
	}
}
