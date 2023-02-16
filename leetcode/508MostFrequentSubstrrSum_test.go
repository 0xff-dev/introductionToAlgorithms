package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindFrequentTreeSum(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 0},
			Right: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 3},
			},
		},
		Right: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 6},
		},
	}
	exp := []int{6}
	if r := findFrequentTreeSum(tree); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	tree = &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: -3},
	}
	exp = []int{-3, 2, 4}
	r := findFrequentTreeSum(tree)
	sort.Ints(r)
	if !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
