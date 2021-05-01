package leetcode

import "testing"

func TestRightSideView(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	r := rightSideView(tree)
	t.Logf("%v", r)

	tree = &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 3,
		},
	}

	r = rightSideView(tree)
	t.Logf("%v", r)

	tree = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	r = rightSideView(tree)
	t.Logf("%v", r)
}
