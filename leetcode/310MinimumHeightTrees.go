package leetcode

/*
超时了
func findMinHeightTrees(n int, edges [][]int) []int {
	dis := make([][]int, n)

	for i := 0; i < n; i++ {
		dis[i] = make([]int, n)
	}
	// 任何一个点到另一个点的路径是唯一的, 所以他俩的距离是可以计算
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		dis[from][to] = 1
		dis[to][from] = 1
	}

	var bfs func(int)
	bfs = func(root int) {
		v := make([]bool, n)
		v[root] = true
		queue := []int{root}
		d := 0
		for len(queue) > 0 {
			nq := make([]int, 0)
			for _, cur := range queue {
				dis[root][cur] = d
				dis[cur][root] = d

				for next, dd := range dis[cur] {
					if dd == 1 && !v[next] {
						//点可以直接到达
						v[next] = true
						nq = append(nq, next)
					}
				}
			}
			queue = nq
			d++
		}
	}
	hm := make(map[int][]int)
	ans := -1
	for root := 0; root < n; root++ {
		bfs(root)
		maxh := 0

		for i := 0; i < n; i++ {
			if root == i {
				continue
			}
			maxh = max(dis[root][i], maxh)
		}
		if _, ok := hm[maxh]; !ok {
			hm[maxh] = make([]int, 0)
		}
		hm[maxh] = append(hm[maxh], root)
		if ans == -1 || maxh < ans {
			ans = maxh
		}
	}
	return hm[ans]
}
*/

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	adj := make([]map[int]struct{}, n)
	for i := 0; i < n; i++ {
		adj[i] = make(map[int]struct{})
	}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		adj[a][b] = struct{}{}
		adj[b][a] = struct{}{}
	}
	l := []int{}
	for i := 0; i < n; i++ {
		if len(adj[i]) == 1 {
			l = append(l, i)
		}
	}
	r := n
	for r > 2 {
		r -= len(l)
		nl := []int{}
		for _, lf := range l {
			var nei int
			for n := range adj[lf] {
				nei = n
			}
			delete(adj[nei], lf)
			if len(adj[nei]) == 1 {
				nl = append(nl, nei)
			}
		}
		l = nl
	}
	return l
}
