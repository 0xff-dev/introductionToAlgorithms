package leetcode

import (
	"reflect"
	"testing"
)

func TestCountSubTrees(t *testing.T) {
	n, edges, labels := 7, [][]int{
		{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6},
	}, "abaedcd"
	exp := []int{2, 1, 1, 1, 1, 1, 1}
	if r := countSubTrees(n, edges, labels); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	n, edges, labels = 4, [][]int{
		{0, 1}, {1, 2}, {0, 3},
	}, "bbbb"
	exp = []int{4, 2, 1, 1}
	if r := countSubTrees(n, edges, labels); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	n, edges, labels = 5, [][]int{
		{0, 1}, {0, 2}, {1, 3}, {0, 4},
	}, "aabab"
	exp = []int{3, 2, 1, 1, 1}
	if r := countSubTrees(n, edges, labels); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
