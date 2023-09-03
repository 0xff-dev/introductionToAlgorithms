package leetcode

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val < val {
		return &TreeNode{Val: val, Left: root}
	}
	// 同时维持着右侧比左侧大的逻辑
	r := root.Right
	if r == nil {
		root.Right = &TreeNode{Val: val}
		return root
	}
	root.Right = insertIntoMaxTree(root.Right, val)
	return root

}
