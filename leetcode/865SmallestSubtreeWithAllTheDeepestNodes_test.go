package leetcode

import "testing"

func TestSubtreeWithAllDeepest(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 4},
			},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 8},
		},
	}
	if r := subtreeWithAllDeepest(tree); r.Val != 2 {
		t.Fatalf("expect 2 get %d", r.Val)
	}

	tree = &TreeNode{
		Val: 1,
	}
	if r := subtreeWithAllDeepest(tree); r.Val != 1 {
		t.Fatalf("expect 2 get %d", r.Val)
	}

	tree = &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{Val: 3},
	}

	if r := subtreeWithAllDeepest(tree); r.Val != 2 {
		t.Fatalf("expect 2 get %d", r.Val)
	}
}
