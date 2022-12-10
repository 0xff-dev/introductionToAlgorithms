package leetcode

import "testing"

func TestMaxProduct1339(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 6},
		},
	}
	if r := maxProduct1339(tree); r != 110 {
		t.Fatalf("expect 110 get %d", r)
	}

	tree = &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
		},
	}
	if r := maxProduct1339(tree); r != 90 {
		t.Fatalf("expect 90 get %d", r)
	}
}
