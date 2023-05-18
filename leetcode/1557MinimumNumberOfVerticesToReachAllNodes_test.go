package leetcode

import (
	"reflect"
	"testing"
)

func TestFindSmallestSetOfVertices(t *testing.T) {
	n := 6
	edges := [][]int{
		{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2},
	}
	exp := []int{0, 3}
	if r := findSmallestSetOfVertices(n, edges); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	n = 5
	edges = [][]int{
		{0, 1}, {2, 1}, {3, 1}, {1, 4}, {2, 4},
	}
	exp = []int{0, 2, 3}
	if r := findSmallestSetOfVertices(n, edges); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
