package leetcode

import "testing"

func TestBinaryTreePaths(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	r := binaryTreePaths(tree)
	t.Logf("%v", r)

	tree = &TreeNode{
		Val: 1,
	}
	r = binaryTreePaths(tree)
	t.Logf("%v", r)
}
