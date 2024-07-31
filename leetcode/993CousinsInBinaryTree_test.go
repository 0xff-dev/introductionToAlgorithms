package leetcode

import "testing"

func TestIsCousins(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 4},
		},
		Right: &TreeNode{Val: 3},
	}
	x, y := 4, 3
	exp := false
	if r := isCousins(tree, x, y); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
	tree = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 5},
		},
	}
	x, y = 5, 4
	exp = true
	if r := isCousins(tree, x, y); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
	tree = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{Val: 3},
	}
	x, y = 2, 3
	exp = false
	if r := isCousins(tree, x, y); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
