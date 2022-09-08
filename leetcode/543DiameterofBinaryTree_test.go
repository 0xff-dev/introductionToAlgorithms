package leetcode

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3},
	}

	if r := diameterOfBinaryTree(tree); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	tree = &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}
	if r := diameterOfBinaryTree(tree); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
