package leetcode

import (
	"reflect"
	"testing"
)

func TestCreateBinaryTree(t *testing.T) {
	descriptions := [][]int{
		{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1},
	}
	exp := &TreeNode{
		Val: 50,
		Left: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 17},
		},
		Right: &TreeNode{
			Val:  80,
			Left: &TreeNode{Val: 19},
		},
	}
	if r := createBinaryTree(descriptions); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	descriptions = [][]int{
		{1, 2, 1}, {2, 3, 0}, {3, 4, 1},
	}
	exp = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 4},
			},
		},
	}
	if r := createBinaryTree(descriptions); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
