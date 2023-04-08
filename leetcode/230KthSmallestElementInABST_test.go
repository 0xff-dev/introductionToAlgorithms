package leetcode

import "testing"

func TestKthSmallet(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{Val: 4},
	}
	if r := kthSmallest230(tree, 1); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	tree = &TreeNode{
		Val:   5,
		Right: &TreeNode{Val: 6},
		Left: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 4},
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 1},
			},
		},
	}
	if r := kthSmallest230(tree, 3); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}
