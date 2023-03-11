package leetcode

import "testing"

func TestAverageOfSubtree(t *testing.T) {
	tree := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val:   5,
			Right: &TreeNode{Val: 6},
		},
	}
	if r := averageOfSubtree(tree); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}

	tree = &TreeNode{Val: 1}
	if r := averageOfSubtree(tree); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
