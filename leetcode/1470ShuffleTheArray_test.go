package leetcode

import (
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	nums, n := []int{2, 5, 1, 3, 4, 7}, 3
	exp := []int{2, 3, 5, 4, 1, 7}
	if r := shuffle(nums, n); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
