package leetcode

import "sort"

type edge struct {
	x, y int
}
type UnionFind struct {
	Father []int
}

func (uf *UnionFind) FindFather(x int) int {
	if uf.Father[x] != x {
		uf.Father[x] = uf.FindFather(uf.Father[x])
	}
	return uf.Father[x]
}

func (uf *UnionFind) Union(x, y int) {
	x = uf.Father[x]
	y = uf.Father[y]
	if x < y {
		uf.Father[y] = x
	} else {
		uf.Father[x] = y
	}
}

const max2421 = 30001

func numberOfGoodPaths(vals []int, edges [][]int) int {
	n := len(vals)
	uf := UnionFind{Father: make([]int, max2421)}
	for i := 0; i < n; i++ {
		uf.Father[i] = i
	}
	idx2node := make(map[int][]int)
	for i, v := range vals {
		if _, ok := idx2node[v]; !ok {
			idx2node[v] = make([]int, 0)
		}
		idx2node[v] = append(idx2node[v], i)
	}

	canReach := make(map[int][]edge)
	for _, e := range edges {
		x, y := e[0], e[1]
		if vals[x] < vals[y] {
			x, y = y, x
		}
		// 分数分数较大的那个
		canReach[vals[x]] = append(canReach[vals[x]], edge{x: x, y: y})
	}
	ans := 0
	sort.Ints(vals)
	pre := -1
	for i := 0; i < n; i++ {
		if pre != -1 && vals[i] == vals[pre] {
			continue
		}
		pre = i
		for _, e := range canReach[vals[i]] {
			if uf.FindFather(e.x) != uf.FindFather(e.y) {
				uf.Union(e.x, e.y)
			}
		}

		count := make(map[int]int)
		for _, node := range idx2node[vals[i]] {
			root := uf.FindFather(node)
			count[root]++
		}
		for _, c := range count {
			ans += c * (c - 1) / 2
		}
	}

	return ans + n
}
