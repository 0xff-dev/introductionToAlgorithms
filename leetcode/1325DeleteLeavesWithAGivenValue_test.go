package leetcode

import (
	"reflect"
	"testing"
)

func TestRemoveLeafNodes(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 2},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 4},
		},
	}
	exp := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 4},
		},
	}

	if r := removeLeafNodes(tree, 2); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	tree = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{Val: 3},
	}
	exp = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 2},
		},
	}
	if r := removeLeafNodes(tree, 3); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	tree = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 2,
				},
			},
		},
	}
	exp = &TreeNode{
		Val: 1,
	}
	if r := removeLeafNodes(tree, 2); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
