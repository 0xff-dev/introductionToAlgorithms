package leetcode

import (
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	exp := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 5,
						Right: &TreeNode{
							Val: 6,
						},
					},
				},
			},
		},
	}
	flatten(tree)
	if !reflect.DeepEqual(exp, tree) {
		t.Fatalf("expect %v get %v", tree, exp)
	}
}
