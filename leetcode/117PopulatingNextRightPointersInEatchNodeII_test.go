package leetcode

import (
	"reflect"
	"testing"
)

func TestConnect117(t *testing.T) {
	tree := &Node1{
		Val: 1,
		Left: &Node1{
			Val:   2,
			Left:  &Node1{Val: 4},
			Right: &Node1{Val: 5},
		},
		Right: &Node1{
			Val:   3,
			Right: &Node1{Val: 7},
		},
	}
	nodes := [6]*Node1{
		{Val: 1},
		{Val: 2},
		{Val: 3},
		{Val: 4},
		{Val: 5},
		{Val: 7},
	}
	nodes[0].Left = nodes[1]
	nodes[0].Right = nodes[2]
	nodes[1].Left = nodes[3]
	nodes[1].Right = nodes[4]
	nodes[2].Right = nodes[5]
	nodes[1].Next = nodes[2]
	nodes[3].Next = nodes[4]
	nodes[4].Next = nodes[5]
	if r := connect117(tree); !reflect.DeepEqual(nodes[0], r) {
		t.Fatalf("expect %v get %v", nodes[0], r)
	}
}
