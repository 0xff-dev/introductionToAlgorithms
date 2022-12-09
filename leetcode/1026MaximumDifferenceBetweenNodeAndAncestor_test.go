package leetcode

import "testing"

func TestMaxAncestorDiff(t *testing.T) {
	tree := &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{
				Val:   6,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 7},
			},
		},
		Right: &TreeNode{
			Val: 10,
			Right: &TreeNode{
				Val: 14,
				Left: &TreeNode{
					Val: 13,
				},
			},
		},
	}

	if r := maxAncestorDiff(tree); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}

	tree = &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 3,
				},
			},
		},
	}

	if r := maxAncestorDiff(tree); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	tree = &TreeNode{
		Val: 2,
		Right: &TreeNode{
			Val: 0,
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
		},
	}
	if r := maxAncestorDiff(tree); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}
