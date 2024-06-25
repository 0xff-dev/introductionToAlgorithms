package leetcode

import "testing"

func TestGetDirections(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 4},
		},
	}
	s, d := 3, 6
	exp := "UURL"
	if r := getDirections(tree, s, d); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
}
