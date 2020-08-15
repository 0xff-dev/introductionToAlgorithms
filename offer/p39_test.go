package offer

import "testing"

func TestBinaryTreeDepth(t *testing.T) {
	tree := &TreeNode{
		Val:   10,
		Left:  &TreeNode{
			Val:   6,
			Left:  &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   13,
			Left:  nil,
			Right: nil,
		},
	}
	if depth := BinaryTreeDepth(tree); depth != 3 {
		t.Fatalf("binary tree depth is %d expect %d", depth, 3)
	}
	tree = nil
	if depth := BinaryTreeDepth(tree); depth != 0 {
		t.Fatalf("binary tree is nil depth shoule be 0")
	}
}

func TestAdvanceIsBalanceBinaryTree(t *testing.T) {
	tree := &TreeNode{
		Val:   10,
		Left:  &TreeNode{
			Val:   6,
			Left:  &TreeNode{
				Val:   1,
				Left:  &TreeNode{
					Val: 4,
				},
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   13,
			Left:  nil,
			Right: nil,
		},
	}
	depth := 0
	isBalanced := AdvanceIsBalanceBinaryTree(tree, &depth)
	if isBalanced {
		t.Fatalf("not balanced")
	}
}