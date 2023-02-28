package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindDuplicateSubtrees(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{Val: 4},
		},
	}

	exp := []int{2, 4}
	r := findDuplicateSubtrees(tree)
	x := make([]int, 0)
	for _, i := range r {
		x = append(x, i.Val)
	}
	sort.Ints(x)
	if !reflect.DeepEqual(exp, x) {
		t.Fatalf("expect %v get %v", exp, x)
	}
}
