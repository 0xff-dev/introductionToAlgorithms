package leetcode

import (
	"reflect"
	"testing"
)

func TestRelativeSortArray(t *testing.T) {
	arr1, arr2 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}
	ans := []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19}
	r := relativeSortArray(arr1, arr2)
	if !reflect.DeepEqual(ans, r) {
		t.Fatalf("expect %v get %v", ans, r)
	}

	arr1, arr2 = []int{28, 6, 22, 8, 44, 17}, []int{22, 28, 8, 6}
	ans = []int{22, 28, 8, 6, 17, 44}
	r = relativeSortArray(arr1, arr2)
	if !reflect.DeepEqual(ans, r) {
		t.Fatalf("expect %v get %v", ans, r)
	}
}
