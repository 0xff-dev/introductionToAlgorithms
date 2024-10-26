package leetcode

import (
	"reflect"
	"testing"
)

func TestTreeQueries(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 2},
		},
		Right: &TreeNode{
			Val:  4,
			Left: &TreeNode{Val: 6},
			Right: &TreeNode{
				Val:   5,
				Right: &TreeNode{Val: 7},
			},
		},
	}
	queries := []int{4}
	exp := []int{2}
	if r := treeQueries(root, queries); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

}
