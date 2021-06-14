package leetcode

import "testing"

func TestIsSubTree(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{Val: 5},
	}

	sub := &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}

	if !isSubtree(tree, sub) {
		t.Fatalf("expect true get false")
	}

	tree.Left.Right.Left = &TreeNode{Val: 0}
	if isSubtree(tree, sub) {
		t.Fatalf("expect false get true")
	}
}
