package leetcode

import "testing"

func TestConstructor1261(t *testing.T) {
	tree := &TreeNode{
		Val:   -1,
		Right: &TreeNode{Val: -1},
	}
	a := Constructor1261(tree)
	if a.Find(1) {
		t.Fatalf("expect false get true")
	}
	if !a.Find(2) {
		t.Fatalf("expect true get false")
	}
}
