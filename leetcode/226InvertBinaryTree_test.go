package leetcode

import "testing"

func TestInvertTree(t *testing.T) {
	tree := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
		Right: &TreeNode{Val: 3, Left: nil, Right: nil},
	}

	newTree := invertTree(tree)
	newTree.Floor()

	tree = &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
			Right: &TreeNode{Val: 3, Left: nil, Right: nil},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 6, Left: nil, Right: nil},
			Right: &TreeNode{Val: 9, Left: nil, Right: nil},
		},
	}

	newTree = invertTree(tree)
	newTree.Floor()
}
