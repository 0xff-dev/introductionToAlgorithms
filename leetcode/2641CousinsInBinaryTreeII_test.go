package leetcode

import (
	"reflect"
	"testing"
)

func TestReplaceValueInTree(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 10},
		},
		Right: &TreeNode{
			Val:   9,
			Right: &TreeNode{Val: 7},
		},
	}
	exp := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val:   0,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: 7},
		},
		Right: &TreeNode{
			Val:   0,
			Right: &TreeNode{Val: 11},
		},
	}
	if r := replaceValueInTree(tree); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	tree = &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}
	exp = &TreeNode{
		Val:   0,
		Left:  &TreeNode{Val: 0},
		Right: &TreeNode{Val: 0},
	}
	if r := replaceValueInTree(tree); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
