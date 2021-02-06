package leetcode

import "testing"

func TestLevelOrder(t *testing.T) {
	tree := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	r := levelOrder(tree)
	t.Log(r)
}
