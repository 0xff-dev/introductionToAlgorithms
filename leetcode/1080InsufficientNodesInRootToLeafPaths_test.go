package leetcode

import (
	"reflect"
	"testing"
)

func TestSufficientSubset(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 8,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
			Right: &TreeNode{
				Val:   -99,
				Left:  &TreeNode{Val: -99},
				Right: &TreeNode{Val: -99},
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   -99,
				Left:  &TreeNode{Val: 12},
				Right: &TreeNode{Val: 13},
			},
			Right: &TreeNode{
				Val:   7,
				Left:  &TreeNode{Val: -99},
				Right: &TreeNode{Val: 14},
			},
		},
	}
	exp := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 8},
				Right: &TreeNode{Val: 9},
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val:   7,
				Right: &TreeNode{Val: 14},
			},
		},
	}
	if r := sufficientSubset(tree, 1); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
