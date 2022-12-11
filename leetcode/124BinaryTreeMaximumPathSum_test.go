package leetcode

import "testing"

func TestMaxPathSum(t *testing.T) {
	tree := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}

	if r := maxPathSum(tree); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}

	tree = &TreeNode{
		Val:  -10,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	if r := maxPathSum(tree); r != 42 {
		t.Fatalf("expect 42 get %d", r)
	}
}
