package leetcode

import "testing"

func TestDistributeCoins(t *testing.T) {
	tree := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 0},
		Right: &TreeNode{Val: 0},
	}
	if r := distributeCoins(tree); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	tree = &TreeNode{
		Val:   0,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 0},
	}
	if r := distributeCoins(tree); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}
