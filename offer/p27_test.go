package offer

import "testing"

func TestConvertBinarySearchTree(t *testing.T) {
	tree := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 14,
			Left: &TreeNode{
				Val:   12,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   16,
				Left:  nil,
				Right: nil,
			},
		},
	}
	ConvertBinarySearchTree(tree)
}
