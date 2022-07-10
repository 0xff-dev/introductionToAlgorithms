package leetcode

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	n1, n2 := []int{1, 2, 2, 1}, []int{2, 2}
	expect := []int{2}
	if r := intersection(n1, n2); !reflect.DeepEqual(expect, r) {
		t.Fatalf("expect %v get %v", expect, r)
	}

	n1, n2 = []int{4, 9, 5}, []int{9, 4, 9, 8, 4}
	expect = []int{9, 4}
	if r := intersection(n1, n2); !reflect.DeepEqual(expect, r) {
		t.Fatalf("expect %v get %v", expect, r)
	}
}
