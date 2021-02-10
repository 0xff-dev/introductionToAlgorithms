package leetcode

import "testing"

func TestZigzagLevelOrder(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{15, nil, nil},
			Right: &TreeNode{7, nil, nil},
		},
	}
	r := zigzagLevelOrder(tree)
	t.Log(r)

	tree = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: &TreeNode{5, nil, nil},
		},
	}
	r = zigzagLevelOrder(tree)
	t.Log(r)
}
