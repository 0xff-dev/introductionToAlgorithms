package leetcode

import (
	"reflect"
	"testing"
)

func TestAllPossibleFBT(t *testing.T) {
	exp := []*TreeNode{
		&TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val:   0,
				Right: &TreeNode{Val: 0},
				Left:  &TreeNode{Val: 0},
			},
		},
		&TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val:   0,
				Right: &TreeNode{Val: 0},
				Left:  &TreeNode{Val: 0},
			},
			Right: &TreeNode{
				Val: 0,
			},
		},
	}
	if r := allPossibleFBT(5); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	exp = []*TreeNode{
		&TreeNode{
			Val:  0,
			Left: &TreeNode{Val: 0},
			Right: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val:   0,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 0},
				},
			},
		},
		&TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val:   0,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 0},
				},
				Right: &TreeNode{Val: 0},
			},
		},
		&TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val:   0,
				Left:  &TreeNode{Val: 0},
				Right: &TreeNode{Val: 0},
			},
			Right: &TreeNode{
				Val:   0,
				Left:  &TreeNode{Val: 0},
				Right: &TreeNode{Val: 0},
			},
		},
		&TreeNode{
			Val: 0,
			Right: &TreeNode{
				Val: 0,
			},
			Left: &TreeNode{
				Val:  0,
				Left: &TreeNode{Val: 0},
				Right: &TreeNode{
					Val:   0,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 0},
				},
			},
		},
		&TreeNode{
			Val:   0,
			Right: &TreeNode{Val: 0},
			Left: &TreeNode{
				Val:   0,
				Right: &TreeNode{Val: 0},
				Left: &TreeNode{
					Val:   0,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 0},
				},
			},
		},
	}
	if r := allPossibleFBT(7); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
