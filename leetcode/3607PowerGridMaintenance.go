package leetcode

import (
	"container/heap"
)

type unionFind3607 struct {
	father []int
}

func (u *unionFind3607) find(x int) int {
	if u.father[x] != x {
		u.father[x] = u.find(u.father[x])
	}
	return u.father[x]
}

func (u *unionFind3607) union(x, y int) {
	fx := u.find(x)
	fy := u.find(y)
	if fx < fy {
		u.father[fy] = fx
	} else {
		u.father[fx] = fy
	}
}

type heapItem3607 struct {
	// state 0表示正常，1表示离线
	c, state, index int
}

type heapList3607 []*heapItem3607

func (h *heapList3607) Len() int {
	return len(*h)
}

func (h *heapList3607) Less(i, j int) bool {
	a, b := (*h)[i], (*h)[j]
	if a.state == b.state {
		return a.c < b.c
	}
	if a.state == 0 {
		return true
	}
	return false
}

func (h *heapList3607) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
	(*h)[i].index = i
	(*h)[j].index = j
}

func (h *heapList3607) Push(x any) {
	item := x.(*heapItem3607)
	item.index = len(*h)
	*h = append(*h, item)
}

func (h *heapList3607) Pop() any {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

type heapManager struct {
	h     *heapList3607
	cache map[int]*heapItem3607
}

func processQueries3607(c int, connections [][]int, queries [][]int) []int {
	// 先通过并查集
	ret := make([]int, 0)
	u := unionFind3607{father: make([]int, c+1)}
	for i := 1; i <= c; i++ {
		u.father[i] = i
	}
	for _, conn := range connections {
		u.union(conn[0], conn[1])
	}
	cache := make(map[int]*heapManager)
	for i := 1; i <= c; i++ {
		f := u.find(i)
		_, ok := cache[f]
		if !ok {
			cache[f] = &heapManager{
				h:     &heapList3607{},
				cache: make(map[int]*heapItem3607),
			}
		}
		item := &heapItem3607{
			c: i, state: 0,
		}
		cache[f].cache[i] = item
		heap.Push(cache[f].h, item)
	}
	for _, q := range queries {
		f := u.find(q[1])
		h := cache[f]
		item := h.cache[q[1]]

		if q[0] == 1 {
			if item.state == 0 {
				// 自己
				ret = append(ret, q[1])
				continue
			}
			top := (*h.h)[0]
			if top.state == 1 {
				ret = append(ret, -1)
				continue
			}
			ret = append(ret, top.c)
			continue
		}

		item.state = 1
		heap.Fix(h.h, item.index)
	}

	return ret
}
