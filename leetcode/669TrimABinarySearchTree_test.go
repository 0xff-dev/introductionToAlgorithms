package leetcode

import (
	"reflect"
	"testing"
)

func TestTrimBST(t *testing.T) {
	tree := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 0},
		Right: &TreeNode{Val: 2},
	}

	exp := &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2},
	}
	if r := trimBST(tree, 1, 2); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
