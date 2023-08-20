package leetcode

import (
	"reflect"
	"testing"
)

func TestSortItems(t *testing.T) {
	n, m := 8, 2
	group := []int{-1, -1, 1, 0, 0, 1, 0, -1}
	beforeItems := [][]int{
		{}, {6}, {5}, {6}, {3, 6}, {}, {}, {},
	}
	exp := []int{6, 3, 4, 5, 2, 0, 7, 1}
	if r := sortItems(n, m, group, beforeItems); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	n, m = 8, 2
	group = []int{-1, -1, 1, 0, 0, 1, 0, -1}
	beforeItems = [][]int{
		{}, {6}, {5}, {6}, {3}, {}, {4}, {},
	}
	if ans := sortItems(n, m, group, beforeItems); ans != nil {
		t.Fatalf("expect nil get %v", ans)
	}
}
