package leetcode

import "testing"

func TestMinDepth(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	if depth := minDepth(tree); depth != 2 {
		t.Fatalf("expect 2 get %d", depth)
	}

	tree = &TreeNode{
		Val: 2,
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
		},
	}

	if depth := minDepth(tree); depth != 5 {
		t.Fatalf("expect 5 get %d", depth)
	}
}
