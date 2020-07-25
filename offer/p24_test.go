package offer

import "testing"

func TestPostOrderRecursive(t *testing.T) {
	tree := &TreeNode{
		Val:   8,
		Left:  &TreeNode{6, &TreeNode{5, nil, nil}, &TreeNode{7, nil, nil}},
		Right: &TreeNode{10, &TreeNode{9, nil, nil}, &TreeNode{11, nil, nil}},
	}
	PostOrderRecursive(tree)
}

func TestPostOrderStack(t *testing.T) {
	tree := &TreeNode{
		Val:   8,
		Left:  &TreeNode{6, &TreeNode{5, nil, nil}, &TreeNode{7, nil, nil}},
		Right: &TreeNode{10, &TreeNode{9, nil, nil}, &TreeNode{11, nil, nil}},
	}
	PostOrderStack(tree)
}
