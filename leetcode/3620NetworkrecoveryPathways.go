package leetcode

import (
	"container/heap"
	"math"
)

type Pair3620 struct {
	First  int
	Second int64
}

type MinHeap3620 []Pair3620

func (h MinHeap3620) Len() int            { return len(h) }
func (h MinHeap3620) Less(i, j int) bool  { return h[i].Second < h[j].Second }
func (h MinHeap3620) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap3620) Push(x interface{}) { *h = append(*h, x.(Pair3620)) }
func (h *MinHeap3620) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findMaxPathScore(edges [][]int, online []bool, k int64) int {
	n := len(online)
	g := make([][]Pair3620, n)
	l, r := int(1e9), 0

	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		if !online[u] || !online[v] {
			continue
		}
		g[u] = append(g[u], Pair3620{v, int64(w)})
		if w < l {
			l = w
		}
		if w > r {
			r = w
		}
	}

	check := func(mid int) bool {
		dis := make([]int64, n)
		for i := range dis {
			dis[i] = math.MaxInt64
		}
		h := &MinHeap3620{}
		heap.Init(h)

		dis[0] = 0
		heap.Push(h, Pair3620{0, 0})

		for h.Len() > 0 {
			top := heap.Pop(h).(Pair3620)
			d, u := top.Second, top.First

			if d > k {
				return false
			}
			if u == n-1 {
				return true
			}
			if d > dis[u] {
				continue
			}

			for _, edge := range g[u] {
				v, w := edge.First, edge.Second
				if w < int64(mid) {
					continue
				}
				if dis[v] > dis[u]+w {
					dis[v] = dis[u] + w
					heap.Push(h, Pair3620{v, dis[v]})
				}
			}
		}
		return false
	}

	if !check(l) {
		return -1
	}

	for l <= r {
		mid := (l + r) >> 1
		if check(mid) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}
