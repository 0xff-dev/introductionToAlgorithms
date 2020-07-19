package offer

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// for test
func (node *TreeNode) PortOrder() {
	if node == nil {
		return
	}
	if node.Left != nil {
		node.Left.PortOrder()
	}
	if node.Right != nil {
		node.Right.PortOrder()
	}
	fmt.Println("val: ", node.Val)
}

// 根据中序，后序构造同理
func Construct(pre, in []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	node := &TreeNode{
		Val:   pre[0],
		Left:  nil,
		Right: nil,
	}
	inOrderIndex := 0
	for ; inOrderIndex < len(in); inOrderIndex++ {
		if in[inOrderIndex] == pre[0] {
			break
		}
	}
	if inOrderIndex > 0 {
		node.Left = Construct(pre[1:inOrderIndex+1], in[:inOrderIndex])
	}
	if inOrderIndex < len(in)-1 {
		node.Right = Construct(pre[inOrderIndex+1:], in[inOrderIndex+1:])
	}
	return node
}
