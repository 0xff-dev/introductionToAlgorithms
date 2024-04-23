package leetcode

import (
	"reflect"
	"testing"
)

func TestFindMinHeightTrees(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 0}, {1, 2}, {1, 3},
	}
	exp := []int{1}
	if r := findMinHeightTrees(n, edges); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	n = 6
	edges = [][]int{
		{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4},
	}
	exp = []int{3, 4}
	if r := findMinHeightTrees(n, edges); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
