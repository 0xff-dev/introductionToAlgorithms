package leetcode

import (
	"testing"
)

func TestSearchBST(t *testing.T) {
	tree := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 7},
	}

	n := searchBST(tree, 2)
	n.Floor()
}
