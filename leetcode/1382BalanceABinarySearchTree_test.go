package leetcode

import (
	"reflect"
	"testing"
)

func TestBalanceBST(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
	}
	newTree := balanceBST(tree)
	exp := &TreeNode{
		Val:  2,
		Left: &TreeNode{Val: 1},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	if !reflect.DeepEqual(exp, newTree) {
		t.Fatalf("expect %v get %v", exp, newTree)
	}
}
