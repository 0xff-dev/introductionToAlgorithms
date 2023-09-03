package leetcode

import (
	"reflect"
	"testing"
)

func TestInsertIntoMaxTree(t *testing.T) {
	tree := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 2},
		},
	}
	exp := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 2},
			},
		},
	}
	if r := insertIntoMaxTree(tree, 5); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	tree = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{Val: 4},
	}
	exp = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 3}},
	}
	if r := insertIntoMaxTree(tree, 3); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	tree = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	exp = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val:  4,
			Left: &TreeNode{Val: 3},
		},
	}
	if r := insertIntoMaxTree(tree, 4); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
