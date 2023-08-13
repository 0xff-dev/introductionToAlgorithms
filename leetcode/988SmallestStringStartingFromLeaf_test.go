package leetcode

import "testing"

func TestSmallestFromLeaf(t *testing.T) {
	tree := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	if r := smallestFromLeaf(tree); r != "dba" {
		t.Fatalf("expect dba get %s", r)
	}

	tree = &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:  0,
			Left: &TreeNode{Val: 1},
		},
		Right: &TreeNode{Val: 1},
	}
	if r := smallestFromLeaf(tree); r != "bae" {
		t.Fatalf("expect bae get %s", r)
	}
}
