package leetcode

import "testing"

func TestLongestZigZag(t *testing.T) {
	t1 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 1},
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val:   1,
						Right: &TreeNode{Val: 1},
					},
				},
			},
		},
	}
	if r := longestZigZag(t1); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	t1 = &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 1},
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 1},
				Left: &TreeNode{
					Val:   1,
					Right: &TreeNode{Val: 1},
				},
			},
		},
	}
	if r := longestZigZag(t1); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}
