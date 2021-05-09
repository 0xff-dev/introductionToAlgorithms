package leetcode

import "testing"

func TestSumRootToLeaf(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
	}

	if r := sumRootToLeaf(tree); r != 22 {
		t.Fatalf("expect 22 get %d", r)
	}

	tree = &TreeNode{
		Val: 1,
	}

	if r := sumRootToLeaf(tree); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
