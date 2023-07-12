package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestEventualSafeNodes(t *testing.T) {
	graph := [][]int{
		{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {},
	}
	exp := []int{2, 4, 5, 6}
	r := eventualSafeNodes(graph)
	sort.Ints(r)
	if !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	graph = [][]int{
		{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {},
	}
	exp = []int{4}
	r = eventualSafeNodes(graph)
	sort.Ints(r)
	if !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	graph = [][]int{
		{}, {0, 2, 3, 4}, {3}, {4}, {},
	}
	exp = []int{0, 1, 2, 3, 4}
	r = eventualSafeNodes(graph)
	sort.Ints(r)
	if !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
