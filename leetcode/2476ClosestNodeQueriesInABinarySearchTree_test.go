package leetcode

import (
	"reflect"
	"testing"
)

func TestClosestNodes(t *testing.T) {
	tree := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:  13,
			Left: &TreeNode{Val: 9},
			Right: &TreeNode{
				Val:  15,
				Left: &TreeNode{Val: 14},
			},
		},
	}
	queries := []int{2, 5, 16}
	exp := [][]int{
		{2, 2}, {4, 6}, {15, -1},
	}
	if r := closestNodes(tree, queries); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	tree = &TreeNode{Val: 4, Right: &TreeNode{Val: 9}}
	queries = []int{3}
	exp = [][]int{{-1, 4}}
	if r := closestNodes(tree, queries); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
