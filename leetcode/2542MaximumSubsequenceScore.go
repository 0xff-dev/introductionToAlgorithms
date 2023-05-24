package leetcode

import (
	"container/heap"
	"sort"
)

/*
func maxScore2542(nums1 []int, nums2 []int, k int) int64 {
	ans := int64(-1)
	maxScore25421(nums1, nums2, 0, k, 0, -1, &ans)
	return ans
}

func maxScore25421(nums1, nums2 []int, index, k int, sum, min int64, ans *int64) {
	if k == 0 {
		if *ans == -1 || sum*min > *ans {
			*ans = sum * min
		}
		return
	}
	if index >= len(nums1) {
		if k == 0 {
			if *ans == -1 || sum*min > *ans {
				*ans = sum * min
			}
		}
		return
	}
	maxScore25421(nums1, nums2, index+1, k, sum, min, ans)
	m := min
	if m == -1 || int64(nums2[index]) < m {
		m = int64(nums2[index])
	}
	maxScore25421(nums1, nums2, index+1, k-1, sum+int64(nums1[index]), m, ans)
}
*/

type pair2542 struct {
	a, b int
}

type topk2542 []int

func (t *topk2542) Len() int {
	return len(*t)
}

func (t *topk2542) Less(i, j int) bool {
	return (*t)[i] < (*t)[j]
}

func (t *topk2542) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}

func (t *topk2542) Push(x interface{}) {
	*t = append(*t, x.(int))
}

func (t *topk2542) Pop() interface{} {
	old := *t
	l := len(old)
	x := old[l-1]
	*t = old[:l-1]
	return x
}
func maxScore2542(nums1, nums2 []int, k int) int64 {
	pairs := make([]pair2542, len(nums1))
	for i := 0; i < len(nums1); i++ {
		pairs[i] = pair2542{nums1[i], nums2[i]}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].b > pairs[j].b
	})
	topKsum := int64(0)
	topKHeap := topk2542([]int{})
	for i := 0; i < k; i++ {
		topKsum += int64(pairs[i].a)
		heap.Push(&topKHeap, pairs[i].a)
	}
	ans := topKsum * int64(pairs[k-1].b)
	for i := k; i < len(nums1); i++ {
		topKsum += int64(pairs[i].a - heap.Pop(&topKHeap).(int))
		heap.Push(&topKHeap, pairs[i].a)
		if r := topKsum * int64(pairs[i].b); r > ans {
			ans = r
		}
	}
	return ans
}
