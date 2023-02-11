package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestDistanceK(t *testing.T) {
	node5 := &TreeNode{
		Val:  5,
		Left: &TreeNode{Val: 6},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: 4},
		},
	}
	tree := &TreeNode{
		Val:  3,
		Left: node5,
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 8},
		},
	}
	exp := []int{1, 4, 7}
	x := distanceK(tree, node5, 2)
	sort.Ints(x)
	if !reflect.DeepEqual(exp, x) {
		t.Fatalf("expect %v get %v", exp, x)
	}
}
