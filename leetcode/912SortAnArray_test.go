package leetcode

import (
	"reflect"
	"testing"
)

func TestSortArray(t *testing.T) {
	n := []int{5, 3, 2, 1}
	exp := []int{1, 2, 3, 5}
	if r := sortArray(n); !reflect.DeepEqual(r, exp) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	n = []int{0, 1}
	exp = []int{0, 1}
	if r := sortArray(n); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	n = []int{0}
	exp = []int{0}
	if r := sortArray(n); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	n = []int{1, 0}
	exp = []int{0, 1}
	if r := sortArray(n); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
