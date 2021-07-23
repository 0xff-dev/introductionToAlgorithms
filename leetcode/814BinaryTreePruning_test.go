package leetcode

import "testing"

func TestPruneTree(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 0,
		},
	}
	t1 := pruneTree(tree)
	t.Logf("t1 == nil: %v", t1 == nil)
	t1.Floor()

	tree = &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
	}
	t1 = pruneTree(tree)
	t.Logf("t1 == nil: %v", t1 == nil)
	t1.Floor()

	tree = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 0,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
	}

	t1 = pruneTree(tree)
	t.Logf("t1 == nil: %v", t1 == nil)
	t1.Floor()
}
