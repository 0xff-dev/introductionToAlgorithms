package leetcode

import "testing"

func TestConstruct449(t *testing.T) {
	tree := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	c := Constructor449()
	s := c.serialize(tree)
	t.Logf("%s", s)
	tt := c.deserialize(s)
	tt.Floor()
}
