package leetcode

func sumRootToLeaf(root *TreeNode) int {
	sum, path := 0, 0
	sumRootToLeafAux(root, path, &sum)
	return sum
}

func sumRootToLeafAux(root *TreeNode, path int, sum *int) {
	if root == nil {
		return
	}

	path <<= 1
	if root.Val == 1 {
		path++
	}

	if root.Left == nil && root.Right == nil {
		*sum += path
		return
	}
	sumRootToLeafAux(root.Left, path, sum)
	sumRootToLeafAux(root.Right, path, sum)
}
