package leetcode

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	nums, k := []int{1, 1, 1, 2, 2, 3}, 2
	if r := topKFrequent(nums, k); !reflect.DeepEqual(r, []int{1, 2}) {
		t.Fatalf("expect %v get %v", []int{1, 2}, r)
	}
}
