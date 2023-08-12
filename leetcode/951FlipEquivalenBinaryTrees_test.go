package leetcode

import "testing"

func TestFlipEquiv(t *testing.T) {
	t1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 4},
			Right: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 8},
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 6},
		},
	}
	t2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 6},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 8},
				Right: &TreeNode{Val: 7},
			},
		},
	}
	if !flipEquiv(t1, t2) {
		t.Fatalf("expect true get flase")
	}

	if !flipEquiv(nil, nil) {
		t.Fatalf("expect true get false")
	}

	t2 = &TreeNode{Val: 1}
	if flipEquiv(nil, t2) {
		t.Fatalf("expect false get true")
	}
}
