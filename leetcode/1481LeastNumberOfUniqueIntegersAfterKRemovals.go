package leetcode

import (
	"container/heap"
)

type eleCount struct {
	n, c int
}

type eleCountList []eleCount

func (e *eleCountList) Len() int {
	return len(*e)
}

func (e *eleCountList) Less(i, j int) bool {
	return (*e)[i].c < (*e)[j].c
}
func (e *eleCountList) Swap(i, j int) {
	(*e)[i], (*e)[j] = (*e)[j], (*e)[i]
}

func (e *eleCountList) Push(x interface{}) {
	*e = append(*e, x.(eleCount))
}

func (e *eleCountList) Pop() interface{} {
	old := *e
	l := len(old)
	x := old[l-1]
	*e = old[:l-1]
	return x
}
func findLeastNumOfUniqueInts(arr []int, k int) int {
	// 4: 1
	// 3:3
	// 2:1
	// 1:2
	nodeCache := make(map[int]int)
	list := eleCountList{}
	for _, n := range arr {
		idx, ok := nodeCache[n]
		if !ok {
			list = append(list, eleCount{n: n, c: 1})
			p := len(list) - 1
			nodeCache[n] = p
			continue
		}
		list[idx].c++
	}
	heap.Init(&list)
	add := 0
	for k > 0 && list.Len() > 0 {
		top := heap.Pop(&list).(eleCount)
		if top.c >= k {
			if top.c != k {
				add = 1
			}
			break
		}
		k -= top.c
	}
	return list.Len() + add
}
