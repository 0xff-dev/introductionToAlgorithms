package leetcode

import (
	"reflect"
	"testing"
)

func TestCircularPermutation(t *testing.T) {
	n, start := 2, 3
	exp := []int{3, 2, 0, 1}
	if r := circularPermutation(n, start); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	n, start = 3, 2
	exp = []int{2, 6, 7, 5, 4, 0, 1, 3}
	if r := circularPermutation(n, start); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
