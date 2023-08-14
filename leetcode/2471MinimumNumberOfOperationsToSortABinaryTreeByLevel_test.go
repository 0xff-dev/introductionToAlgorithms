package leetcode

import "testing"

func TestMinimumOperations2471(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: 6},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:  8,
				Left: &TreeNode{Val: 9},
			},
			Right: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 10},
			},
		},
	}
	if r := minimumOperations2471(tree); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}
