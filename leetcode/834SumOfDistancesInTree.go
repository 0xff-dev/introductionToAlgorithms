package leetcode

/*
func sumOfDistancesInTree(n int, edges [][]int) []int {
	ans := make([]int, n)
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, n)
	}
	for _, e := range edges {
		// 互联, 单项
		adj[e[0]][e[1]] = 1
		adj[e[1]][e[0]] = 1
	}

	var dfs func(int, int, int)
	dfs = func(from, next, deepth int) {
		if from >= n {
			return
		}
		for idx, rel := range adj[next] {
			if idx == from || next == idx {
				continue
			}
			if rel > 0 {
				zero := adj[from][idx] == 0
				if adj[from][idx] == 0 || rel+deepth < adj[from][idx] {
					adj[from][idx] = rel + deepth
					adj[idx][from] = adj[from][idx]
				}
				if zero {
					dfs(from, idx, deepth+1)
				}
			}
				if rel > 1 && (adj[from][idx] == 0 || rel+deepth < adj[from][idx]) {
					adj[from][idx] = rel+deepth
					adj[idx][from] = adj[from][idx]
					dfs(from, idx, deepth+1)
				}
		}
	}
	for u := 0; u < n; u++ {
		for idx, rel := range adj[u] {
			if rel == 1 {
				dfs(u, idx, 1)
			}
		}
	}

	for u := 0; u < n; u++ {
		s := 0
		for rel := 0; rel < n; rel++ {
			if rel == u {
				continue
			}
			s += adj[u][rel]
			if adj[u][rel] == 0 {
				s += adj[rel][u]
			}
		}
		ans[u] = s
	}
	return ans
}
*/

func sumOfDistancesInTree(n int, edges [][]int) []int {
	adj := make(map[int]map[int]struct{})
	for u := 0; u < n; u++ {
		adj[u] = make(map[int]struct{})
	}
	for _, e := range edges {
		adj[e[0]][e[1]] = struct{}{}
		adj[e[1]][e[0]] = struct{}{}
	}

	ans := make([]int, n)
	count := make([]int, n)
	for i := 0; i < n; i++ {
		count[i] = 1
	}
	var dfs func(int, int)
	var dfs1 func(int, int)
	dfs = func(node, parent int) {
		for child := range adj[node] {
			if child != parent {
				dfs(child, node)
				count[node] += count[child]
				ans[node] += ans[child] + count[child]
			}
		}
	}
	dfs1 = func(node, parent int) {
		for child := range adj[node] {
			if child != parent {
				ans[child] = ans[node] - count[child] + n - count[child]
				dfs1(child, node)
			}
		}
	}
	dfs(0, -1)
	dfs1(0, -1)
	return ans
}
