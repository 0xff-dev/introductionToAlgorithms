package leetcode

import "testing"

func TestIsCompleteTree(t *testing.T) {
	t1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
		},
	}
	if !isCompleteTree(t1) {
		t.Fatalf("expect true get false")
	}

	t1 = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 7},
		},
	}
	if isCompleteTree(t1) {
		t.Fatalf("expect false get true")
	}
}
