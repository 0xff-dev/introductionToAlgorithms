package leetcode

import "testing"

func TestPseudoPalindromicPaths(t *testing.T) {
	tree := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 1},
		},
		Right: &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 1},
		},
	}
	if r := pseudoPalindromicPaths(tree); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	tree = &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 1}},
		},
		Right: &TreeNode{Val: 1},
	}
	if r := pseudoPalindromicPaths(tree); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	tree = &TreeNode{Val: 1}
	if r := pseudoPalindromicPaths(tree); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
