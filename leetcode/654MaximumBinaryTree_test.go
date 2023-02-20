package leetcode

import (
	"reflect"
	"testing"
)

func TestConstructMaximumBinaryTree(t *testing.T) {
	nums := []int{3, 2, 1, 6, 0, 5}
	tree := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 1},
			},
		},
		Right: &TreeNode{
			Val:  5,
			Left: &TreeNode{Val: 0},
		},
	}
	if r := constructMaximumBinaryTree(nums); !reflect.DeepEqual(r, tree) {
		t.Fatalf("expect %v get %v", tree, r)
	}
}
