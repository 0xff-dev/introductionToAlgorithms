package leetcode

import (
	"reflect"
	"testing"
)

func TestFlatten430(t *testing.T) {
	node1 := &FNode{Val: 1}
	node2 := &FNode{Val: 2}
	node3 := &FNode{Val: 3}

	node1.Child = node3
	node1.Next = node2
	node2.Prev = node1

	node11 := &FNode{Val: 1}
	node22 := &FNode{Val: 2}
	node33 := &FNode{Val: 3}
	node11.Next = node33
	node33.Prev = node11
	node33.Next = node22
	node22.Prev = node33
	r := flatten430(node1)
	if !reflect.DeepEqual(node11, r) {
		t.Fatalf("expect %v get %v", node11, r)
	}
}
