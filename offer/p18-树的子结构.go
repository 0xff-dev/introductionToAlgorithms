package offer

func IsChildren(a, b *TreeNode) bool {
	subTree := false
	if a != nil && b != nil {
		if a.Val == b.Val {
			subTree = hasSameStructure(a, b)
		}
		if !subTree {
			subTree = IsChildren(a.Left, b)
		}
		if !subTree {
			subTree = IsChildren(a.Right, b)
		}
	}
	return subTree
}

func hasSameStructure(a, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil {
		return false
	}
	return hasSameStructure(a.Left, b.Left) && hasSameStructure(a.Right, b.Right)
}
