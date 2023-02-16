package leetcode

import "testing"

func TestFindBottomLeftValue(t *testing.T) {
	tree := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}
	if r := findBottomLeftValue(tree); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	tree = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 7,
				},
			},
			Right: &TreeNode{Val: 6},
		},
	}
	if r := findBottomLeftValue(tree); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}
}
