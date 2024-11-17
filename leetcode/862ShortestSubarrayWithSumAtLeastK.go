package leetcode

import (
	"container/list"
	"math"
)

func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	prefixSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	minLength := math.MaxInt32
	deque := list.New()

	for i := 0; i <= n; i++ {
		for deque.Len() > 0 && prefixSum[i]-prefixSum[deque.Front().Value.(int)] >= k {
			minLength = min(minLength, i-deque.Remove(deque.Front()).(int))
		}

		for deque.Len() > 0 && prefixSum[i] <= prefixSum[deque.Back().Value.(int)] {
			deque.Remove(deque.Back())
		}

		deque.PushBack(i)
	}

	if minLength == math.MaxInt32 {
		return -1
	}
	return minLength
}
