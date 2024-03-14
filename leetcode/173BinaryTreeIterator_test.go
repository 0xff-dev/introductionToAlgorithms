package leetcode

import "testing"

func TestConstructor173(t *testing.T) {
	tree := &TreeNode{
		Val:  7,
		Left: &TreeNode{Val: 3},
		Right: &TreeNode{
			Val:   15,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 20},
		},
	}
	c := Constructor173(tree)
	if r := c.Next(); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	if r := c.Next(); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}
	if !c.HasNext() {
		t.Fatalf("expect true get false")
	}

	if r := c.Next(); r != 9 {
		t.Fatalf("expect 9 get %d", r)
	}
	if !c.HasNext() {
		t.Fatalf("expect true get false")
	}
	if r := c.Next(); r != 15 {
		t.Fatalf("expect 15 get %d", r)
	}
	if !c.HasNext() {
		t.Fatalf("expect true get false")
	}
	if r := c.Next(); r != 20 {
		t.Fatalf("expect 7 get %d", r)
	}
	if c.HasNext() {
		t.Fatalf("expect false get true")
	}

}
