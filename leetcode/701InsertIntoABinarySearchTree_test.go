package leetcode

import (
	"reflect"
	"testing"
)

func TestInsertIntoBST(t *testing.T) {
	tree := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 7},
	}
	exp := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 5}},
	}

	if r := insertIntoBST(tree, 5); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	tree = &TreeNode{
		Val: 40,
		Left: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 10},
			Right: &TreeNode{Val: 30},
		},
		Right: &TreeNode{
			Val:   60,
			Left:  &TreeNode{Val: 50},
			Right: &TreeNode{Val: 70},
		},
	}
	exp = &TreeNode{
		Val: 40,
		Left: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 10},
			Right: &TreeNode{Val: 30, Left: &TreeNode{Val: 25}},
		},
		Right: &TreeNode{
			Val:   60,
			Left:  &TreeNode{Val: 50},
			Right: &TreeNode{Val: 70},
		},
	}
	if r := insertIntoBST(tree, 25); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
