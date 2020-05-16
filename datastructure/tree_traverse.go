/*
	Achieve binary tree, binary search tree, read black tree, balanced binary tree
*/

package datastructure

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
