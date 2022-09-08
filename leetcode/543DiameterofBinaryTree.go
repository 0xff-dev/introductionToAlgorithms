package leetcode

import "fmt"

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	diameterOfBinaryTreeHelper(root, &ans)
	return ans
}

func diameterOfBinaryTreeHelper(root *TreeNode, ans *int) (int, int) {
	if root == nil {
		return 0, 0
	}

	ll, lr := diameterOfBinaryTreeHelper(root.Left, ans)
	fmt.Printf("[%d] ll: %d, lr: %d\n", root.Val, ll, lr)
	rl, rr := diameterOfBinaryTreeHelper(root.Right, ans)
	fmt.Printf("[%d] rl: %d, rr: %d\n", root.Val, rl, rr)

	left, right := ll, rl
	if left < lr {
		left = lr
	}
	if right < rr {
		right = rr
	}
	//left, right = left+1, right+1
	fmt.Printf("accross [%d] exp: %d\n", root.Val, left+right)
	if left+right > *ans {
		*ans = left + right
	}
	return left + 1, right + 1
}
