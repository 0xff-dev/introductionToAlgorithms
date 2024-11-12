package leetcode

import (
	"reflect"
	"testing"
)

func TestMaximumBeauty(t *testing.T) {
	items, queries, exp := [][]int{{1, 2}, {3, 2}, {2, 4}, {5, 6}, {3, 5}}, []int{1, 2, 3, 4, 5, 6}, []int{2, 4, 5, 5, 6, 6}
	if r := maximumBeauty(items, queries); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	items, queries, exp = [][]int{{1, 2}, {1, 2}, {1, 3}, {1, 4}}, []int{1}, []int{4}
	if r := maximumBeauty(items, queries); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	items, queries, exp = [][]int{{10, 1000}}, []int{5}, []int{0}
	if r := maximumBeauty(items, queries); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
