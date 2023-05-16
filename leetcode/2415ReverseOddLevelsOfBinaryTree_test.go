package leetcode

import (
	"reflect"
	"testing"
)

func TestReverseOddLevels(t *testing.T) {
	tree := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 8},
			Right: &TreeNode{Val: 13},
		},
		Right: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 21},
			Right: &TreeNode{Val: 34},
		},
	}
	expTree := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 8},
			Right: &TreeNode{Val: 13},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 21},
			Right: &TreeNode{Val: 34},
		},
	}
	tt := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 8},
			Right: &TreeNode{Val: 13},
		},
		Right: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 21},
			Right: &TreeNode{Val: 34},
		},
	}
	if r := reverseOddLevels(tree); !reflect.DeepEqual(expTree, r) {
		t.Fatalf("expect %v get %v", expTree, r)
	}
	if r := reverseOddLevels1(tree); !reflect.DeepEqual(r, tt) {
		t.Fatalf("expect %v get %v", tt, r)
	}

	tree = &TreeNode{
		Val:   7,
		Left:  &TreeNode{Val: 13},
		Right: &TreeNode{Val: 11},
	}
	expTree = &TreeNode{
		Val:   7,
		Left:  &TreeNode{Val: 11},
		Right: &TreeNode{Val: 13},
	}
	tt = &TreeNode{
		Val:   7,
		Left:  &TreeNode{Val: 13},
		Right: &TreeNode{Val: 11},
	}
	if r := reverseOddLevels(tree); !reflect.DeepEqual(expTree, r) {
		t.Fatalf("expect %v get %v", expTree, r)
	}
	if r := reverseOddLevels1(tree); !reflect.DeepEqual(r, tt) {
		t.Fatalf("expect %v get %v", tt, r)
	}
}
