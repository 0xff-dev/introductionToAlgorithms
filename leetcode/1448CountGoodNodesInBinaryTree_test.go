package leetcode

import "testing"

func TestGoodNodes(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 5},
		},
	}
	if r := goodNodes(tree); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	tree = &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 2},
		},
	}
	if r := goodNodes(tree); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	tree = &TreeNode{Val: 1}
	if r := goodNodes(tree); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	tree = &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  5,
			Left: &TreeNode{Val: 4},
		},
	}
	if r := goodNodes(tree); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	tree = &TreeNode{
		Val: 9,
		Right: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 6},
		},
	}
	if r := goodNodes(tree); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
