package leetcode

import (
	"testing"
)

func TestRangeSumBST(t *testing.T) {
	tree := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 7},
		},
		Right: &TreeNode{
			Val: 15,
			Right: &TreeNode{
				Val: 18,
			},
		},
	}
	if r := rangeSumBST(tree, 7, 15); r != 32 {
		t.Fatalf("expect 32 get %d", r)
	}

	tree = &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val:  7,
				Left: &TreeNode{Val: 6},
			},
		},
		Right: &TreeNode{
			Val:   15,
			Left:  &TreeNode{Val: 13},
			Right: &TreeNode{Val: 18},
		},
	}

	if r := rangeSumBST(tree, 6, 10); r != 23 {
		t.Fatalf("expect 23 get %d", r)
	}
}
