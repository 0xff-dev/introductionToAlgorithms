package leetcode

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	nodeValues := make([]int, 0)
	isValidBSTAux(root, &nodeValues)
	return isSorted(nodeValues)
}

func isValidBSTAux(root *TreeNode, nodeValues *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		isValidBSTAux(root.Left, nodeValues)
	}
	*nodeValues = append(*nodeValues, root.Val)
	if root.Right != nil {
		isValidBSTAux(root.Right, nodeValues)
	}
}

func isSorted(nodeValues []int) bool {
	n := len(nodeValues)
	for idx := n - 1; idx > 0; idx-- {
		if nodeValues[idx] <= nodeValues[idx-1] {
			return false
		}
	}
	return true
}
