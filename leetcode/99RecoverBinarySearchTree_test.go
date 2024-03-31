package leetcode

import (
	"reflect"
	"testing"
)

func TestRecoverTree99(t *testing.T) {
	tree := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 3, Right: &TreeNode{Val: 2}},
	}
	exp := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
	}
	recoverTree99(tree)
	if !reflect.DeepEqual(exp, tree) {
		t.Fatalf("expect %v get %v", exp, tree)
	}

	tree = &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 2}},
	}
	exp = &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}},
	}
	recoverTree99(tree)
	if !reflect.DeepEqual(exp, tree) {
		t.Fatalf("expect %v get %v", exp, tree)
	}
}
