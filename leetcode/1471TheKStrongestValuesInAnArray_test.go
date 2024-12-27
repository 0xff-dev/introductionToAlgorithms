package leetcode

import (
	"reflect"
	"testing"
)

func TestGetStrongest(t *testing.T) {
	arr, k, exp := []int{1, 2, 3, 4, 5}, 2, []int{5, 1}
	if r := getStrongest(arr, k); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	arr, k, exp = []int{1, 1, 3, 5, 5}, 2, []int{5, 5}
	if r := getStrongest(arr, k); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	arr, k, exp = []int{6, 7, 11, 7, 6, 8}, 5, []int{11, 8, 6, 6, 7}
	if r := getStrongest(arr, k); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
