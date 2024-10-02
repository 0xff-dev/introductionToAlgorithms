package leetcode

import (
	"reflect"
	"testing"
)

func TestArrayRankTransform(t *testing.T) {
	arr, exp := []int{40, 10, 20, 30}, []int{4, 1, 2, 3}
	if r := arrayRankTransform(arr); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	arr, exp = []int{100, 100, 100}, []int{1, 1, 1}
	if r := arrayRankTransform(arr); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	arr, exp = []int{37, 12, 28, 9, 100, 56, 80, 5, 12}, []int{5, 3, 4, 2, 8, 6, 7, 1, 3}
	if r := arrayRankTransform(arr); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
