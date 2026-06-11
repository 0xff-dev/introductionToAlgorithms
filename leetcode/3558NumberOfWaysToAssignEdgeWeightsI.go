package leetcode

const mod3558 = 1000000007

func qpow3558(base, exp, mod int64) int64 {
	var res int64 = 1
	base = base % mod
	for exp > 0 {
		if exp&1 == 1 {
			res = (res * base) % mod
		}
		base = (base * base) % mod
		exp >>= 1
	}
	return res
}

func assignEdgeWeights(edges [][]int) int {
	n := len(edges) + 1

	adj := make([][]int, n+1)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	type pair struct {
		node  int
		depth int
	}

	queue := make([]pair, 0, n)
	visited := make([]bool, n+1)

	queue = append(queue, pair{node: 1, depth: 0})
	visited[1] = true

	maxDepth := 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.depth > maxDepth {
			maxDepth = curr.depth
		}

		for _, neighbor := range adj[curr.node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, pair{node: neighbor, depth: curr.depth + 1})
			}
		}
	}

	L := maxDepth
	if L == 0 {
		return 0
	}

	ans := qpow3558(2, int64(L-1), int64(mod3558))
	return int(ans)
}
