/*
	Achieve binary tree, binary search tree, read black tree, balanced binary tree
*/

package datastructure

import "fmt"

// traverse
func (tNode *treeNode) PreOrder() {
	if tNode == nil {
		return
	}
	tNode.Print()
	if tNode.Left != nil {
		tNode.Left.PreOrder()
	}
	if tNode.Right != nil {
		tNode.Right.PreOrder()
	}
}

func (tNode *treeNode) InOrder() {
	if tNode == nil {
		return
	}
	if tNode.Left != nil {
		tNode.Left.InOrder()
	}
	tNode.Print()
	if tNode.Right != nil {
		tNode.Right.InOrder()
	}
}

func (tNode *treeNode) PostOrder() {
	if tNode == nil {
		return
	}
	if tNode.Left != nil {
		tNode.Left.PostOrder()
	}
	if tNode.Right != nil {
		tNode.Right.PostOrder()
	}
	tNode.Print()
}

// stack
func (tNode *treeNode) PreOrder1() {
	if tNode == nil {
		return
	}
	stack := NewStack()
	root := tNode
	for root != nil || !stack.Empty() {
		for ; root != nil; root = root.Left {
			root.Print()
			stack.Push(root)
		}
		if !stack.Empty() {
			root, _ = stack.Pop()
			root = root.Right
		}
	}
}

func (tNode *treeNode) InOrder1() {
	if tNode == nil {
		return
	}
	stack := NewStack()
	root := tNode
	for root != nil || !stack.Empty() {
		for ; root != nil; root = root.Left {
			stack.Push(root)
		}
		if !stack.Empty() {
			root, _ = stack.Pop()
			root.Print()
			root = root.Right
		}
	}
}

func (tNode *treeNode) PostOrder1() {
	if tNode == nil {
		return
	}
	stack := NewStack()
	root := tNode
	for root != nil || !stack.Empty() {
		for ; root != nil; root = root.Left {
			stack.Push(root)
			stack.Push(root)
		}
		if !stack.Empty() {
			root, _ = stack.Pop()
			if topNod, err := stack.Top(); err == nil && topNod == root {
				root = root.Right
			} else {
				root.Print()
				root = nil
			}
		}
	}
}

func (tNode *treeNode) Floor() {
	if tNode == nil {
		return
	}
	queue := NewQueue()
	queue.EnQueue(tNode)
	for !queue.Empty() {
		node := queue.DeQueue()
		node.Print()
		if node.Left != nil {
			queue.EnQueue(node.Left)
		}
		if node.Right != nil {
			queue.EnQueue(node.Right)
		}
	}
}

// 其他的树的常用算法总结
// 所有的节点数量
func (tNode *treeNode) AllNodeQuantity() int {
	if tNode == nil {
		return 0
	}
	return tNode.Left.AllNodeQuantity() + tNode.Right.AllNodeQuantity() + 1
}

// 叶子节点数量
func (tNode *treeNode) LeavesQuantity() int {
	if tNode == nil {
		return 0
	}
	if tNode.Right == nil && tNode.Left == nil {
		return 1
	}
	return tNode.Left.LeavesQuantity() + tNode.Right.LeavesQuantity()
}

//树的深度
func (tNode *treeNode) TreeDepth() int {
	if tNode == nil {
		return 0
	}
	left := tNode.Left.TreeDepth() + 1
	right := tNode.Right.TreeDepth() + 1
	if left > right {
		return left
	}
	return right
}

// 第k层的节点数量
func (tNode *treeNode) KthNodeQuantity(k int) int {
	if tNode == nil {
		return 0
	}
	if k == 1 {
		return 1
	}
	return tNode.Left.KthNodeQuantity(k-1) + tNode.Right.KthNodeQuantity(k-1)
}

// pre, in ==> post
func FindPostOrder(pre, in []int) {
	if len(pre) == 0 {
		return
	}
	if len(pre) == 1 {
		fmt.Println("item: ", pre[0])
		return
	}
	root := pre[0]
	rootIndex := 0
	for ; in[rootIndex] != root; rootIndex++ {
	}
	if rootIndex > 0 {
		FindPostOrder(pre[1:rootIndex+1], in[:rootIndex])
	}
	if rootIndex < len(in)-1 {
		FindPostOrder(pre[rootIndex+1:], in[rootIndex+1:])
	}
	fmt.Println("item: ", root)
}
