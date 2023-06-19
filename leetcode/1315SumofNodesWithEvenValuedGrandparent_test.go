package leetcode

import "testing"

func TestSumEvenGrandparent(t *testing.T) {
	tree := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 9},
			},
			Right: &TreeNode{
				Val:   7,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 4},
			},
		},
		Right: &TreeNode{
			Val:  8,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{
				Val:   3,
				Right: &TreeNode{Val: 5},
			},
		},
	}
	if r := sumEvenGrandparent(tree); r != 18 {
		t.Fatalf("expect 18 get %d", r)
	}

	tree = &TreeNode{Val: 1}
	if r := sumEvenGrandparent(tree); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}
