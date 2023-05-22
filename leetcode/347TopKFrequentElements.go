package leetcode

import (
	"container/heap"
	"sort"
)

type fq struct {
	v, c int
}

// 大根堆，
// quick select
func topKFrequent(nums []int, k int) []int {
	numCount := make(map[int]int)
	for _, n := range nums {
		numCount[n]++
	}
	sortObj := make([]fq, 0)
	for k, v := range numCount {
		sortObj = append(sortObj, fq{k, v})
	}
	sort.Slice(sortObj, func(i, j int) bool {
		return sortObj[i].c < sortObj[j].c
	})

	result := make([]int, 0)
	for idx := len(sortObj) - 1; idx >= 0 && k > 0; idx, k = idx-1, k-1 {
		result = append(result, sortObj[idx].v)
	}
	return result
}

type fqs []fq

func (f *fqs) Len() int {
	return len(*f)
}

func (f *fqs) Less(i, j int) bool {
	return (*f)[i].c > (*f)[j].c
}

func (f *fqs) Swap(i, j int) {
	(*f)[i], (*f)[j] = (*f)[j], (*f)[i]
}

func (f *fqs) Push(x interface{}) {
	*f = append(*f, x.(fq))
}

func (f *fqs) Pop() interface{} {
	old := *f
	l := len(old)
	x := old[l-1]
	*f = old[:l-1]
	return x
}

func topKFrequent1(nums []int, k int) []int {
	keys := make(map[int]*fq)
	for _, n := range nums {
		if _, ok := keys[n]; !ok {
			keys[n] = &fq{v: n}
		}
		keys[n].c++
	}
	h := fqs([]fq{})
	for _, f := range keys {
		heap.Push(&h, *f)
	}
	ans := make([]int, k)
	index := 0
	for k > 0 && h.Len() > 0 {
		x := heap.Pop(&h).(fq)
		ans[index] = x.v
		index++
		k--
	}
	return ans
}
