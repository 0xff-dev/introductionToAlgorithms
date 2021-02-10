package leetcode

import "testing"

func TestSumNumbers(t *testing.T) {
	tree := &TreeNode{
		Val:   1,
		Left:  &TreeNode{2, nil, nil},
		Right: &TreeNode{3, nil, nil},
	}
	r := sumNumbers(tree)
	if r != 25 {
		t.Fatalf("expect 25 get %d", r)
	}

	tree = &TreeNode{
		Val:   4,
		Left:  &TreeNode{9, &TreeNode{5, nil, nil}, &TreeNode{1, nil, nil}},
		Right: &TreeNode{0, nil, nil},
	}
	if r = sumNumbers(tree); r != 1026 {
		t.Fatalf("expect 1026 get %d", r)
	}

	tree = &TreeNode{
		Val:   4,
		Left:  &TreeNode{9, nil, &TreeNode{1, nil, nil}},
		Right: &TreeNode{0, nil, nil},
	}
	if r = sumNumbers(tree); r != 531 {
		t.Fatalf("expect 531 get %d", r)
	}
}
