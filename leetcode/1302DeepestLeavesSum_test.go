package leetcode

import "testing"

func TestDeepestLeavesSum(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:  4,
				Left: &TreeNode{Val: 7},
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val:   6,
				Right: &TreeNode{Val: 8},
			},
		},
	}

	if r := deepestLeavesSum(tree); r != 15 {
		t.Fatalf("expect 15 get %d", r)
	}
}
