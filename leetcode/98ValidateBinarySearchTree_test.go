package leetcode

import "testing"

func TestIsValidBST(t *testing.T) {
	tree := &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}
	if !isValidBST(tree) {
		t.Fatalf("is bst tree")
	}

	tree = &TreeNode{5, &TreeNode{1, nil, nil}, &TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{6, nil, nil}}}
	if isValidBST(tree) {
		t.Fatalf("isn't bst tree")
	}

	tree = &TreeNode{5, &TreeNode{1, nil, nil}, &TreeNode{6, &TreeNode{3, nil, nil}, &TreeNode{6, nil, nil}}}
	if isValidBST(tree) {
		t.Fatalf("isn't bst tree")
	}

	tree = &TreeNode{1, &TreeNode{1, nil, nil}, nil}
	if isValidBST(tree) {
		t.Fatalf("isn't bst tree")
	}
}
