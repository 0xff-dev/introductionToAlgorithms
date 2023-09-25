package leetcode

/*
type edge743 [][]int

func (e *edge743) Len() int {
	return len(*e)
}

func (e *edge743) Swap(i, j int) {
	(*e)[i], (*e)[j] = (*e)[j], (*e)[i]
}

func (e *edge743) Less(i, j int) bool {
	return (*e)[i][2] < (*e)[j][2]
}

func (e *edge743) Push(x interface{}) {
	*e = append(*e, x.([]int))
}

func (e *edge743) Pop() interface{} {
	old := *e
	l := len(old)
	x := old[l-1]
	*e = old[:l-1]
	return x
}

type unionFind743 struct {
	father []int
}

func (u *unionFind743) Find(x int) int {
	if u.father[x] != x {
		u.father[x] = u.father[u.father[x]]
	}
	return u.father[x]
}

func (u *unionFind743) Union(fx, fy int) {
	if fx < fy {
		u.father[fy] = fx
	} else {
		u.father[fx] = fy
	}
}

func networkDelayTime(times [][]int, n int, k int) int {
	adj := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		adj[i] = make([]int, n+1)
	}
	for _, t := range times {
		from, to, weight := t[0], t[1], t[2]
		adj[from][to] = weight
		adj[to][from] = weight
	}
	edges := edge743{}
	for _, time := range times {
		heap.Push(&edges, time)
	}
	u := unionFind743{father: make([]int, n+1)}
	used := make([]bool, n+1)
	leftNode := n
	for i := 0; i <= n; i++ {
		// 自成一派
		u.father[i] = i
	}
	ans := 0
	for leftNode > 0 {
		if edges.Len() == 0 {
			break
		}
		e := heap.Pop(&edges).([]int)
		fx := u.Find(e[0])
		fy := u.Find(e[1])
		if fx != fy {
			u.Union(fx, fy)
			if !used[fx] {
				used[fx] = true
				leftNode--
			}
			if !used[fy] {
				used[fy] = true
				leftNode--
			}
			ans += e[2]
		}
		// 1, 2
		if leftNode == 0 {
			return ans
		}
	}
	if leftNode == 0 {
		return ans
	}
	return -1
}
*/

// 这道题开始理解为无向图，用了heap+unionfind，结果发现是单向图
func networkDelayTime(times [][]int, n int, k int) int {
	adj := make(map[int]map[int]int)
	for _, t := range times {
		from, to, weigth := t[0], t[1], t[2]
		if _, ok := adj[from]; !ok {
			adj[from] = make(map[int]int)
		}
		adj[from][to] = weigth
	}
	visited := make([]int, n+1)
	for i := 0; i <= n; i++ {
		visited[i] = -1
	}
	visited[k] = 0

	queue := []int{k}
	for len(queue) > 0 {
		nq := make([]int, 0)
		for _, q := range queue {
			for n, w := range adj[q] {
				costTime := visited[q] + w
				if visited[n] == -1 || visited[n] > costTime {
					visited[n] = costTime
					nq = append(nq, n)
				}
			}
		}
		queue = nq
	}
	ans := -1
	for i := 1; i <= n; i++ {
		if visited[i] == -1 {
			return -1
		}
		if ans == -1 || ans < visited[i] {
			ans = visited[i]
		}
	}
	return ans
}
