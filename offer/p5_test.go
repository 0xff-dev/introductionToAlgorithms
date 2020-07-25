package offer

import "testing"

func TestPrintListReversing(t *testing.T) {
	list := &Node{
		Val: 1,
		Next: &Node{
			Val: 2,
			Next: &Node{
				Val:  3,
				Next: nil,
			},
		},
	}
	PrintListReversing(list)
}
