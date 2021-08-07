package leetcode

import "testing"

func TestLevelOrder1(t *testing.T) {
	tree := &MultiChildNode{
		Val: 1,
		Children: []*MultiChildNode{
			&MultiChildNode{
				Val: 3,
				Children: []*MultiChildNode{
					&MultiChildNode{Val: 5},
					&MultiChildNode{Val: 6},
				},
			},
			&MultiChildNode{Val: 2},
			&MultiChildNode{Val: 4},
		},
	}

	r := levelOrder1(tree)
	t.Logf("%v", r)
}
