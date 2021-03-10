package leetcode

import "testing"

func TestDeleteNode(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  6,
			Left: nil,
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	root := deleteNode(tree, 7)
	root.Floor()

	tree = &TreeNode{5, nil, nil}
	if root = deleteNode(tree, 5); root != nil {
		t.Fatalf("expect nil")
	}
}
