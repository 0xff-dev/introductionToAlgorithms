package leetcode

import (
	"reflect"
	"testing"
)

func TestAllPathsSourceTarget(t *testing.T) {
	graph := [][]int{
		{1, 2}, {3}, {3}, {},
	}

	exp := [][]int{
		{0, 1, 3}, {0, 2, 3},
	}
	if r := allPathsSourceTarget(graph); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	graph = [][]int{
		{4, 3, 1}, {3, 2, 4}, {3}, {4}, {},
	}
	exp = [][]int{
		{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4},
	}
	if r := allPathsSourceTarget(graph); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
