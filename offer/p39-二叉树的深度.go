package offer

import "fmt"

func BinaryTreeDepth(tree *TreeNode) int {
	if tree == nil {
		return 0
	}
	left := BinaryTreeDepth(tree.Left) + 1
	right := BinaryTreeDepth(tree.Right) + 1
	if left > right {
		return left
	}
	return right
}

func IsBalanceBinaryTree(tree *TreeNode) bool {
	if tree == nil {
		return true
	}
	left := BinaryTreeDepth(tree.Left)
	right := BinaryTreeDepth(tree.Right)
	difference := left - right
	if difference < -1 || difference > 1 {
		return false
	}
	return IsBalanceBinaryTree(tree.Left) && IsBalanceBinaryTree(tree.Right)
}

func max(args ...int) int {
	if len(args) == 0 {
		return -1
	}
	maxNum := args[0]
	for index := 1; index < len(args); index++ {
		if args[index] > maxNum {
			maxNum = args[index]
		}
	}
	return maxNum
}

func AdvanceIsBalanceBinaryTree(tree *TreeNode, depth *int) bool {
	if tree == nil {
		*depth = 0
		return true
	}
	left, right := 0, 0
	if AdvanceIsBalanceBinaryTree(tree.Left, &left) && AdvanceIsBalanceBinaryTree(tree.Right, &right) {
		diff := left - right
		if diff >= -1 && diff <= 1 {
			fmt.Println("diff: ", diff)
			*depth = max(left, right) + 1
			return true
		}
	}
	return false
}
