package leetcode

import "testing"

func TestFindPath(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{
			Val:   -1,
			Right: &TreeNode{Val: 1},
		},
	}

	if r := findPath(tree, 5); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	if r := findPath2(tree, 5); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	tree = &TreeNode{Val: 1}
	if r := findPath(tree, 1); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	if r := findPath2(tree, 1); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	tree = &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: -2},
			},
			Right: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 1},
			},
		},
		Right: &TreeNode{
			Val:   -3,
			Right: &TreeNode{Val: 11},
		},
	}
	if r := findPath(tree, 8); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	if r := findPath2(tree, 8); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	tree = &TreeNode{Val: 1}
	if r := findPath(tree, 0); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	if r := findPath2(tree, 0); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}
