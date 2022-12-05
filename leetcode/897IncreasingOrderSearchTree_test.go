package leetcode

import (
	"reflect"
	"testing"
)

func TestIncreasingBST(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 4},
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 1},
			},
		},
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Val:   8,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 9},
			},
		},
	}

	exp := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 5,
						Right: &TreeNode{
							Val: 6,
							Right: &TreeNode{
								Val: 7,
								Right: &TreeNode{
									Val: 8,
									Right: &TreeNode{
										Val: 9,
									},
								},
							},
						},
					},
				},
			},
		},
	}
	s := increasingBST(tree)
	if !reflect.DeepEqual(exp, s) {
		t.Fatalf("expect %v get %v", exp, s)
	}

	tree = &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 7},
	}

	exp = &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	s = increasingBST(tree)
	if !reflect.DeepEqual(exp, s) {
		t.Fatalf("expect %v get %v", exp, s)
	}
}
