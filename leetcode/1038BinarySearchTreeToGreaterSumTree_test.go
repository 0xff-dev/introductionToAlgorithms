package leetcode

import (
	"reflect"
	"testing"
)

func TestBstToGst(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 3},
			},
		},
		Right: &TreeNode{
			Val:  6,
			Left: &TreeNode{Val: 5},
			Right: &TreeNode{
				Val:   7,
				Right: &TreeNode{Val: 8},
			},
		},
	}
	exp := &TreeNode{
		Val: 30,
		Left: &TreeNode{
			Val: 36,
			Left: &TreeNode{
				Val: 36,
			},
			Right: &TreeNode{
				Val:   35,
				Right: &TreeNode{Val: 33},
			},
		},
		Right: &TreeNode{
			Val:  21,
			Left: &TreeNode{Val: 26},
			Right: &TreeNode{
				Val:   15,
				Right: &TreeNode{Val: 8},
			},
		},
	}
	if r := bstToGst(root); !reflect.DeepEqual(r, exp) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
