package leetcode

/*
	func buildBST1038(values []int, left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := (right-left)/2 + left
		node := &TreeNode{Val: values[mid]}
		node.Left = buildBST1038(values, left, mid-1)
		node.Right = buildBST1038(values, mid+1, right)
		return node
	}
*/
func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	var preOrder func(*TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		sum += root.Val
		if root.Left != nil {
			preOrder(root.Left)
		}
		if root.Right != nil {
			preOrder(root.Right)
		}
	}
	preOrder(root)
	/*
		tmp := sum
		for i := 0; i < len(nodes); i++ {
			tmp -= nodes[i]
			nodes[i] += tmp
		}
	*/
	tmpSum := 0
	var inOrder func(*TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			inOrder(root.Left)
		}
		tmpSum += root.Val
		root.Val += sum - tmpSum
		if root.Right != nil {
			inOrder(root.Right)
		}
	}
	inOrder(root)
	return root
}
