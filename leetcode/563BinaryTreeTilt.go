package leetcode

func findTilt(root *TreeNode) int {

	r := 0
	findTiltAux(root, &r)
	return r
}

func findTiltAux(root *TreeNode, r *int) int {
	if root == nil {
		return 0
	}

	left := findTiltAux(root.Left, r)
	right := findTiltAux(root.Right, r)
	v := root.Val
	root.Val = left - right
	if root.Val < 0 {
		root.Val = -root.Val
	}
	*r += root.Val
	return left + right + v
}
