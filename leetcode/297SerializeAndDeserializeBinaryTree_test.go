package leetcode

import (
	"reflect"
	"testing"
)

func TestConstructor297(t *testing.T) {
	c := Constructor297()
	tree := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
	}
	str := c.serialize(tree)
	if r := c.deserialize(str); !reflect.DeepEqual(r, tree) {
		r.Floor()
		t.Fatalf("expect %v get %v", tree, r)
	}

	tree = &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2},
	}
	str = c.serialize(tree)
	if r := c.deserialize(str); !reflect.DeepEqual(r, tree) {
		r.Floor()
		t.Fatalf("expect %v get %v", tree, r)
	}

	tree = &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: -7},
		Right: &TreeNode{Val: -3},
	}
	str = c.serialize(tree)
	if r := c.deserialize(str); !reflect.DeepEqual(r, tree) {
		r.Floor()
		t.Fatalf("expect %v get %v", tree, r)
	}
}
