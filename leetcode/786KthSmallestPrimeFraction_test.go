package leetcode

import (
	"reflect"
	"testing"
)

func TestKthSmallestPrimeFraction(t *testing.T) {
	arr := []int{1, 2, 3, 5}
	k := 3
	exp := []int{2, 5}
	if r := kthSmallestPrimeFraction(arr, k); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	arr = []int{1, 7}
	k = 1
	exp = []int{1, 7}
	if r := kthSmallestPrimeFraction(arr, k); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
