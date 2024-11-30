package leetcode

import "container/heap"

type heap3342 [][4]int

func (h *heap3342) Len() int {
	return len(*h)
}

func (h *heap3342) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *heap3342) Less(i, j int) bool {
	return (*h)[i][2] < (*h)[j][2]
}

func (h *heap3342) Push(x any) {
	*h = append(*h, x.([4]int))
}

func (h *heap3342) Pop() any {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

var dirs3342 = [][2]int{
	{0, 1}, {0, -1}, {-1, 0}, {1, 0},
}

func minTimeToReach3342(moveTime [][]int) int {
	m, n := len(moveTime), len(moveTime[0])
	h := heap3342{{0, 0, 0, 1}}
	visited := map[[2]int]struct{}{
		[2]int{0, 0}: {},
	}
	for h.Len() > 0 {
		top := heap.Pop(&h).([4]int)
		if top[0] == m-1 && top[1] == n-1 {
			return top[2]
		}
		for _, dir := range dirs3342 {
			nx, ny := top[0]+dir[0], top[1]+dir[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n {
				if _, ok := visited[[2]int{nx, ny}]; !ok {
					visited[[2]int{nx, ny}] = struct{}{}
					cost := moveTime[nx][ny]
					if cost <= top[2] {
						cost = top[2] + top[3]
					} else {
						cost += top[3]
					}
					heap.Push(&h, [4]int{nx, ny, cost, 3 - top[3]})
				}
			}
		}
	}
	return -1
}
