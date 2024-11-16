package leetcode

import (
	"reflect"
	"testing"
)

func TestResultsArray(t *testing.T) {
	nums, k, exp := []int{1, 2, 3, 4, 3, 2, 5}, 3, []int{3, 4, -1, -1, -1}
	if r := resultsArray(nums, k); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	nums, k, exp = []int{2, 2, 2, 2, 2}, 4, []int{-1, -1}
	if r := resultsArray(nums, k); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	nums, k, exp = []int{3, 2, 3, 2, 3, 2}, 2, []int{-1, 3, -1, 3, -1}
	if r := resultsArray(nums, k); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
