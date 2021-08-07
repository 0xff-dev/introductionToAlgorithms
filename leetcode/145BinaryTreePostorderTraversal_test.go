package leetcode

import "testing"

func TestPostorder(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}

	r := postorderTraversal(tree)
	t.Logf("%v", r)
}
