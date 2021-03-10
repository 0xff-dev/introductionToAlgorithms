package leetcode

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	val := root.Val

	if val > key {
		root.Left = deleteNode(root.Left, key)
	} else if val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left != nil && root.Right != nil {
			// 寻找后继节点
			successorNode := findSuccessorNode(root.Right)
			root.Val = successorNode.Val
			root.Right = deleteNode(root.Right, successorNode.Val)
		} else if root.Left != nil {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}

func findSuccessorNode(root *TreeNode) *TreeNode {
	for ; root.Left != nil; root = root.Left {
	}
	return root
}
