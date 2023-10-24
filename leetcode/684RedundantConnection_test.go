package leetcode

import (
	"reflect"
	"testing"
)

func TestFindRedunantConnection(t *testing.T) {
	edges := [][]int{
		{1, 2}, {1, 3}, {2, 3},
	}
	exp := []int{2, 3}
	if r := findRedundantConnection(edges); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	edges = [][]int{
		{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5},
	}
	exp = []int{1, 4}
	if r := findRedundantConnection(edges); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
