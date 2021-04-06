package leetcode

import "testing"

func TestIsSameTree(t *testing.T) {
	tree1 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{2, nil, nil},
		Right: &TreeNode{3, nil, nil},
	}

	if !isSameTree(tree1, tree1) {
		t.Fatalf("expect true return false")
	}

	if isSameTree(tree1, nil) {
		t.Fatalf("expece false return true")
	}

	tree2 := &TreeNode{1, nil, nil}
	if isSameTree(tree1, tree2) {
		t.Fatalf("expect false return true")
	}

	tree1 = &TreeNode{
		Val:   1,
		Left:  &TreeNode{2, nil, nil},
		Right: &TreeNode{1, nil, nil},
	}
	tree2 = &TreeNode{
		Val:   1,
		Left:  &TreeNode{1, nil, nil},
		Right: &TreeNode{2, nil, nil},
	}

	if isSameTree(tree1, tree2) {
		t.Fatalf("expect false return true")
	}
}
