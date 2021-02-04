package leetcode

import "testing"

func TestInOrderTraversal(t *testing.T) {
	tree := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	r := inorderTraversal(tree)
	t.Log(r)
}
