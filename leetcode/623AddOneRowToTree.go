package leetcode

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		node := &TreeNode{Val: val, Left: root}
		return node
	}
	addOneRow623(root, val, 1, depth)
	return root
}

func addOneRow623(root *TreeNode, val, depth, target int) {
	if root == nil {
		return
	}

	if depth == target-1 {
		// need to add node
		leftNode := &TreeNode{Val: val}
		rightNode := &TreeNode{Val: val}

		originalLeft := root.Left
		originalRight := root.Right

		leftNode.Left = originalLeft
		rightNode.Right = originalRight
		root.Left = leftNode
		root.Right = rightNode
		return
	}

	addOneRow623(root.Left, val, depth+1, target)
	addOneRow623(root.Right, val, depth+1, target)
}
