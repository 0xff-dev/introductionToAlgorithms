package leetcode

import "testing"

func TestNAryPreorder(t *testing.T) {
	tree := &MultiChildNode{
		Val: 1,
		Children: []*MultiChildNode{
			{
				Val: 3,
				Children: []*MultiChildNode{
					{
						Val: 5,
					},
					{
						Val: 6,
					},
				},
			},
			{
				Val: 2,
			},
			{
				Val: 4,
			},
		},
	}

	nodes := nAryPreorder(tree)
	t.Logf("%v", nodes)
}
