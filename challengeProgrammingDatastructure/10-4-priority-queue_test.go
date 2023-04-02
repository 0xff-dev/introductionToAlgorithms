package challengeprogrammingdatastructure

import (
	"sort"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	nums := []int{9, 3, 7, 4, 1, 8, 12, 17}
	my := MyHeap{
		data: make([]int, 0),
		less: func(i, j int) bool {
			return i > j
		},
	}
	for _, n := range nums {
		my.Insert(n)
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	for i, n := range nums {
		x := my.Pop()
		t.Logf("loop: %d, pop %d, expect %d\n", i, x, n)
	}
}
