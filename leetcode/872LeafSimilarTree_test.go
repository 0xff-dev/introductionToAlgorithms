package leetcode

import "testing"

func TestLeafSimilar(t *testing.T) {
	t1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  5,
			Left: &TreeNode{Val: 6},
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 4},
			},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 8},
		},
	}
	t2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7},
		},
		Right: &TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 4},
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 9},
				Right: &TreeNode{Val: 8},
			},
		},
	}

	if !leafSimilar(t1, t2) {
		t.Fatalf("expect true get false")
	}

	t1 = &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	t2 = &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 2},
	}

	if leafSimilar(t1, t2) {
		t.Fatalf("expect false get true")
	}
}
