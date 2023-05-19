package leetcode

type unionFind785 struct {
	father []int
}

func (u *unionFind785) findFather(x int) int {
	if u.father[x] != x {
		u.father[x] = u.findFather(u.father[x])
	}
	return u.father[x]
}
func (u *unionFind785) union(x, y int) {
	fx := u.findFather(x)
	fy := u.findFather(y)
	if fx < fy {
		u.father[y] = fx
	} else {
		u.father[x] = fy
	}
}
func isBipartite(graph [][]int) bool {
	n := len(graph)

	u := unionFind785{father: make([]int, n)}
	for i := 0; i < n; i++ {
		u.father[i] = i
	}

	for from, next := range graph {
		for _, to := range next {
			f1 := u.findFather(from)
			f2 := u.findFather(to)
			u.union(f1, f2)
		}
	}

	nodes := make([]int, n)
	adj := make([][]bool, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]bool, n)
	}

	for from, next := range graph {
		for _, to := range next {
			adj[from][to] = true
			adj[to][from] = true
		}
	}
	distinct := make(map[int]struct{})
	for _, x := range u.father {
		distinct[u.findFather(x)] = struct{}{}
	}
	var dfs func(int, int) bool
	dfs = func(start, color int) bool {
		if nodes[start] != 0 {
			return nodes[start] == color
		}
		nodes[start] = color

		for _, next := range graph[start] {
			if !dfs(next, 3-color) {
				return false
			}
		}
		return true
	}
	for k := range distinct {
		if !dfs(k, 1) {
			return false
		}
	}
	return true
}
