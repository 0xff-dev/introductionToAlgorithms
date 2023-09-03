package leetcode

import "testing"

func TestAmountOfTime(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 10,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	if r := amountOfTime(tree, 3); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	tree = &TreeNode{Val: 1}
	if r := amountOfTime(tree, 1); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}
