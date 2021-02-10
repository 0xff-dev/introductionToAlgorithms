package leetcode

func buildTree1(preorder []int, inorder []int) *TreeNode {
	return buildTree1Aux(preorder, inorder)
}

func buildTree1Aux(preOrder, inOrder []int) *TreeNode {
	if len(preOrder) == 0 {
		return nil
	}
	if len(preOrder) == 1 {
		return &TreeNode{preOrder[0], nil, nil}
	}

	root := preOrder[0]
	node := &TreeNode{root, nil, nil}
	idx := 0
	for ; idx < len(preOrder) && inOrder[idx] != root; idx++ {
	}

	if idx < len(preOrder)-1 {
		// right child
		node.Right = buildTree1Aux(preOrder[idx+1:], inOrder[idx+1:])
	}
	if idx > 0 {
		// left child
		node.Left = buildTree1Aux(preOrder[1:idx+1], inOrder[:idx])
	}

	return node
}
