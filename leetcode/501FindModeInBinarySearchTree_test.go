package leetcode

import (
	"reflect"
	"testing"
)

func TestFindMode(t *testing.T) {
	tree := &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2},
	}
	exp := []int{1, 2}
	if r := findMode(tree); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	tree = &TreeNode{
		Val: 1,
	}
	exp = []int{1}
	if r := findMode(tree); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
