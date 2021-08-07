package leetcode

import "testing"

func TestPreorderTraversal(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3},
		},
	}
	r := preorderTraversal(tree)
	t.Logf("%v", r)
}
