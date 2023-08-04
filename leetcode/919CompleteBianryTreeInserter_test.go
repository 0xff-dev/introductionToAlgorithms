package leetcode

import "testing"

func TestConstructor919(t *testing.T) {
	tree := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}

	o := Constructor919(tree)
	if r := o.Insert(3); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	if r := o.Insert(4); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	x := o.Get_root()
	x.Floor()
}
