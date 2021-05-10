package leetcode

import "fmt"

func binaryTreePaths(root *TreeNode) []string {
	r := make([]string, 0)
	binaryTreePathsAux(root, &r, "")
	return r
}

func binaryTreePathsAux(root *TreeNode, r *[]string, path string) {
	if root == nil {
		return
	}

	rootV := fmt.Sprintf("%d", root.Val)
	if path != "" {
		rootV = "->" + rootV
	}
	path += rootV

	if root.Left == nil && root.Right == nil {
		*r = append(*r, path)
		return
	}

	binaryTreePathsAux(root.Left, r, path)
	binaryTreePathsAux(root.Right, r, path)
}
