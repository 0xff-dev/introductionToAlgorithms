package leetcode

import (
	"reflect"
	"testing"
)

func TestXorQueries(t *testing.T) {
	arr, queries := []int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}
	exp := []int{2, 7, 14, 8}
	if r := xorQueries(arr, queries); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	arr, queries = []int{4, 8, 2, 10}, [][]int{
		{2, 3}, {1, 3}, {0, 0}, {0, 3},
	}
	exp = []int{8, 0, 4, 4}
	if r := xorQueries(arr, queries); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
