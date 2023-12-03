package leetcode

import (
	"reflect"
	"testing"
)

func TestFlipMatchVoyage(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	voyage := []int{2, 1}
	exp := []int{-1}
	if r := flipMatchVoyage(tree, voyage); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	tree = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	voyage = []int{1, 3, 2}
	exp = []int{1}
	if r := flipMatchVoyage(tree, voyage); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	voyage = []int{1, 2, 3}
	exp = []int{}
	if r := flipMatchVoyage(tree, voyage); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
