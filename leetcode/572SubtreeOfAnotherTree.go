package leetcode

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	res := false
	isSub(root, subRoot, &res)
	return res
}

func isSub(root, subRoot *TreeNode, res *bool) {
	if *res {
		return
	}

	*res = isSubtreeAux(root, subRoot)

	if !*res && root != nil {
		isSub(root.Left, subRoot, res)
		isSub(root.Right, subRoot, res)
	}
}
func isSubtreeAux(root, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if root == nil || subRoot == nil {
		return false
	}

	if root.Val != subRoot.Val {
		return false
	}

	return isSubtreeAux(root.Left, subRoot.Left) && isSubtreeAux(root.Right, subRoot.Right)
}
