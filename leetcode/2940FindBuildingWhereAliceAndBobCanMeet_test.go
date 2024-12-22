package leetcode

import (
	"reflect"
	"testing"
)

func TestLeftmostBuildingQueries(t *testing.T) {
	heights, queries, exp := []int{6, 4, 8, 5, 2, 7}, [][]int{{0, 1}, {0, 3}, {2, 4}, {3, 4}, {2, 2}}, []int{2, 5, -1, 5, 2}
	if r := leftmostBuildingQueries(heights, queries); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	heights, queries, exp = []int{5, 3, 8, 2, 6, 1, 4, 6}, [][]int{{0, 7}, {3, 5}, {5, 2}, {3, 0}, {1, 6}}, []int{7, 6, -1, 4, 6}
	if r := leftmostBuildingQueries(heights, queries); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
