package leetcode

import "testing"

func TestFindTilt(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	if r := findTilt(tree); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	tree = &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 9,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	if r := findTilt(tree); r != 15 {
		t.Fatalf("expect 15 get %d", r)
	}
}
