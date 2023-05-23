package leetcode

import (
	"container/heap"
)

type KthLargest struct {
	k    int
	nums []int
}

func (k *KthLargest) Len() int {
	return len(k.nums)
}

func (k *KthLargest) Less(i, j int) bool {
	return k.nums[i] < k.nums[j]
}
func (k *KthLargest) Swap(i, j int) {
	k.nums[i], k.nums[j] = k.nums[j], k.nums[i]
}

func (k *KthLargest) Push(x interface{}) {
	k.nums = append(k.nums, x.(int))
}

func (k *KthLargest) Pop() interface{} {
	old := k.nums
	l := len(k.nums)
	x := old[l-1]
	k.nums = old[:l-1]
	return x
}
func Constructor703(k int, nums []int) KthLargest {
	// 2: 2, 3
	x := KthLargest{k: k, nums: make([]int, 0)}
	for i := 0; i < len(nums); i++ {
		heap.Push(&x, nums[i])
		if i >= k {
			heap.Pop(&x)
		}
	}
	return x
}

// 1
func (this *KthLargest) Add(val int) int {
	//开始为空
	heap.Push(this, val)
	if len(this.nums) > this.k {
		heap.Pop(this)
	}
	x := heap.Pop(this).(int)
	heap.Push(this, x)
	return x
}
