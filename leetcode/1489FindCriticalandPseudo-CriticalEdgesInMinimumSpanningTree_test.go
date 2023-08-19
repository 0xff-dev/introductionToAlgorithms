package leetcode

import (
	"reflect"
	"testing"
)

func TestFindCriticalAndPseudoCriticalEdges(t *testing.T) {
	n := 5
	edges := [][]int{
		{0, 1, 1}, {1, 2, 1}, {2, 3, 2}, {0, 3, 2}, {0, 4, 3}, {3, 4, 3}, {1, 4, 6},
	}
	exp := [][]int{
		{0, 1}, {2, 3, 4, 5},
	}
	if r := findCriticalAndPseudoCriticalEdges(n, edges); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	n = 4
	edges = [][]int{
		{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {0, 3, 1},
	}
	exp = [][]int{
		{}, {0, 1, 2, 3},
	}
	if r := findCriticalAndPseudoCriticalEdges(n, edges); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
