package leetcode

import (
	"fmt"
	"testing"
)

func TestDelNodes(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7},
		},
	}
	ans := delNodes(tree, []int{3, 5})
	for _, te := range ans {
		te.Floor()
		fmt.Println("---")
	}
}
