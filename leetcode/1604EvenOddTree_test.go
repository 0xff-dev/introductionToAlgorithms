package leetcode

import "testing"

func TestIsEvenOddTree(t *testing.T) {
	t1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 12,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 6,
				},
			},
			Right: &TreeNode{
				Val:   9,
				Right: &TreeNode{Val: 2},
			},
		},
	}
	if !isEvenOddTree(t1) {
		t.Fatalf("expect true get false")
	}
	t1 = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 7},
		},
	}
	if isEvenOddTree(t1) {
		t.Fatalf("expect false get true")
	}
	t1 = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 7},
		},
	}
	if isEvenOddTree(t1) {
		t.Fatalf("expect false get true")
	}
}
