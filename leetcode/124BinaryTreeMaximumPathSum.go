package leetcode

func maxPathSum(root *TreeNode) int {
	ans := 0
	init := false
	maxPathSum1241(root, &ans, &init)
	return ans
}

func maxPathSum1241(root *TreeNode, ans *int, ansInit *bool) (int, int) {
	if root == nil {
		return 0, 0
	}

	ll, lr := maxPathSum1241(root.Left, ans, ansInit)
	rl, rr := maxPathSum1241(root.Right, ans, ansInit)
	lmax := ll
	if lr > lmax {
		lmax = lr
	}

	rmax := rr
	if rl > rmax {
		rmax = rl
	}

	returnL, returnR := root.Val, root.Val
	l := lmax + root.Val
	r := rmax + root.Val
	if !*ansInit || l > *ans {
		*ans = l
		*ansInit = true
	}
	if !*ansInit || r > *ans {
		*ans = r
		*ansInit = true
	}
	if r := l + rmax; !*ansInit || r > *ans {
		*ans = r
		*ansInit = true
	}
	if !*ansInit || root.Val > *ans {
		*ans = root.Val
		*ansInit = true
	}
	if l > root.Val {
		returnL = l
	}
	if r > root.Val {
		returnR = r
	}
	return returnL, returnR
}
