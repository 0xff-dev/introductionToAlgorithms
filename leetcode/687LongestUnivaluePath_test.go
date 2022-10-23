package leetcode

import "testing"

func TestLongestUnivaluePath(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 1},
		},
		Right: &TreeNode{
			Val:   5,
			Right: &TreeNode{Val: 5},
		},
	}
	if r := longestUnivaluePath(tree); r != 2 {
		t.Fatalf("expect %d get %d", 2, r)
	}

	tree = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	if r := longestUnivaluePath(tree); r != 2 {
		t.Fatalf("expect %d get %d", 2, r)
	}

	tree = &TreeNode{
		Val:  4,
		Left: &TreeNode{Val: -8},
		Right: &TreeNode{
			Val: -3,
			Right: &TreeNode{
				Val:   -3,
				Right: &TreeNode{Val: 6},
			},
			Left: &TreeNode{
				Val:   -9,
				Left:  &TreeNode{Val: -7},
				Right: &TreeNode{Val: -4},
			},
		},
	}

	if r := longestUnivaluePath(tree); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
