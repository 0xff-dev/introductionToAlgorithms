package leetcode

import "testing"

func TestCheckTree(t *testing.T) {
	tree := &TreeNode{
		Val:   10,
		Left:  &TreeNode{Val: 4},
		Right: &TreeNode{Val: 6},
	}
	if !checkTree(tree) {
		t.Fatalf("expect true get false")
	}

	tree.Right.Val = 5
	if checkTree(tree) {
		t.Fatalf("expect false get true")
	}
}
