package challengeprogrammingdatastructure

import "testing"

func TestHeap(t *testing.T) {
	nums := []int{9, 3, 7, 4, 1, 8, 12, 17}
	nums1 := []int{9, 3, 7, 4, 1, 8, 12, 17}
	maxHeap := func(i, j int) bool {
		return i > j
	}
	minHeap := func(i, j int) bool {
		return i < j
	}
	heapSort(nums, maxHeap)
	t.Logf("nums %v\n", nums)
	heapSort(nums1, minHeap)
	t.Logf("nums1 %v\n", nums1)
}
