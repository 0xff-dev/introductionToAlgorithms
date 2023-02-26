package leetcode

import "testing"

func TestWidthOfBinaryTree(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 6},
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val:  9,
				Left: &TreeNode{Val: 7},
			},
		},
	}
	if r := widthOfBinaryTree(tree); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}
}
