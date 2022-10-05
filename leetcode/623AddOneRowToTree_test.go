package leetcode

import (
	"reflect"
	"testing"
)

func TestAddOneRow(t *testing.T) {
	tree := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 1},
		},
		Right: &TreeNode{
			Val:  6,
			Left: &TreeNode{Val: 5},
		},
	}

	exp := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 1},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val:  6,
				Left: &TreeNode{Val: 5},
			},
		},
	}

	addOneRow(tree, 1, 2)
	if !reflect.DeepEqual(tree, exp) {
		t.Fatalf("expect %v get %v", exp, tree)
	}
}
