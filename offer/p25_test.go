package offer

import "testing"

func TestFindPath(t *testing.T) {
	tree := &TreeNode{
		Val:   10,
		Left:  &TreeNode{
			Val:   5,
			Left:  &TreeNode{4, nil, &TreeNode{3, nil, nil}},
			Right: &TreeNode{7, nil,nil},
		},
		Right: &TreeNode{12, nil, nil},
	}
	FindPath(tree, 22)
}
