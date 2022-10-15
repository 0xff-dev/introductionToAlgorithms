package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindEvenNumbers(t *testing.T) {
	d := []int{2, 1, 3, 0}
	exp := []int{102, 120, 130, 132, 210, 230, 302, 310, 312, 320}
	r := findEvenNumbers(d)
	sort.Ints(r)
	if !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	d = []int{2, 2, 8, 8, 2}
	exp = []int{222, 228, 282, 288, 822, 828, 882}
	r = findEvenNumbers(d)
	sort.Ints(r)

	if !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	d = []int{1, 3, 5}
	exp = []int{}
	r = findEvenNumbers(d)
	sort.Ints(r)

	if !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
