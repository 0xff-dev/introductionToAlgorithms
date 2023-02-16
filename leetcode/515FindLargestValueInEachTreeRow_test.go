package leetcode

import (
	"reflect"
	"testing"
)

func TestLargestValues(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 9},
		},
	}
	exp := []int{1, 3, 9}
	if r := largestValues(tree); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	tree = &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	exp = []int{1, 3}
	if r := largestValues(tree); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
