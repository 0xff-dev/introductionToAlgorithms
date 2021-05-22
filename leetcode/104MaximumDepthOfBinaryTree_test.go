package leetcode

import "testing"

func TestMaxDepth(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
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

	if depth := maxDepth(tree); depth != 3 {
		t.Fatalf("expect 3 get %d", depth)
	}

	tree = &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}

	if depth := maxDepth(tree); depth != 2 {
		t.Fatalf("expect 2 get %d", depth)
	}
}
