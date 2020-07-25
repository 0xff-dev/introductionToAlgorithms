package offer

import "testing"

func TestMirrorBinaryTree(t *testing.T) {
	tree := &TreeNode{
		Val:   8,
		Left:  &TreeNode{6, &TreeNode{5, nil, nil}, &TreeNode{7, nil, nil}},
		Right: &TreeNode{10, &TreeNode{9, nil, nil}, &TreeNode{11, nil, nil}},
	}
	tree.PortOrder()
	MirrorBinaryTree(tree)
	tree.PortOrder()
}
