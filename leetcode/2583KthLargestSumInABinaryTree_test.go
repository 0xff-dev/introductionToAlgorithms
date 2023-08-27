package leetcode

import "testing"

func TestKthLargestLevelSum(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 6},
			},
			Right: &TreeNode{Val: 1},
		},
		Right: &TreeNode{
			Val:   9,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 7},
		},
	}
	if r := kthLargestLevelSum(tree, 2); r != 13 {
		t.Fatalf("expect 13 get %d", r)
	}
}
