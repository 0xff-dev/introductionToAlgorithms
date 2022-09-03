package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestNumSame(t *testing.T) {
	n, k := 3, 7
	exp := []int{181, 292, 707, 818, 929}
	r := numsSameConsecDiff(n, k)
	sort.Ints(exp)
	sort.Ints(r)
	if !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	n, k = 2, 1
	exp = []int{10, 12, 21, 23, 32, 34, 43, 45, 54, 56, 65, 67, 76, 78, 87, 89, 98}
	r = numsSameConsecDiff(n, k)
	sort.Ints(exp)
	sort.Ints(r)

	if !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	n, k = 5, 0
	exp = []int{11111, 22222, 33333, 44444, 55555, 66666, 77777, 88888, 99999}
	r = numsSameConsecDiff(n, k)
	if !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
