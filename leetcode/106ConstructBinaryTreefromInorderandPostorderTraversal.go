package leetcode

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	return buildTreeAux(inorder, postorder)
}

func buildTreeAux(inorder, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	if len(postorder) == 1 {
		return &TreeNode{Val: postorder[0], Left: nil, Right: nil}
	}

	root := postorder[len(postorder)-1]
	node := &TreeNode{
		Val:   root,
		Left:  nil,
		Right: nil,
	}
	idx := 0

	for ; idx < len(inorder) && inorder[idx] != root; idx++ {
	}

	if idx < len(postorder)-1 {
		node.Right = buildTreeAux(inorder[idx+1:], postorder[idx:len(postorder)-1])
	}

	if idx > 0 {
		node.Left = buildTreeAux(inorder[:idx], postorder[:idx])
	}
	return node
}
