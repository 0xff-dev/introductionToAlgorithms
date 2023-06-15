package leetcode

import "testing"

func TestMaxLevelSum(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: -8},
		},
		Right: &TreeNode{Val: 0},
	}
	if r := maxLevelSum(tree); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
