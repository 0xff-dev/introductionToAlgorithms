package leetcode

import "math"

type LCA struct {
	n int
	m int
	d []int
	e [][]int
	f [][]int
}

func NewLCA(edges [][]int, root int) *LCA {
	lca := &LCA{}
	lca.n = len(edges) + 1
	lca.m = int(math.Log2(float64(lca.n))) + 1

	lca.e = make([][]int, lca.n+1)
	lca.d = make([]int, lca.n+1)
	lca.f = make([][]int, lca.n+1)
	for i := 0; i <= lca.n; i++ {
		lca.e[i] = make([]int, 0)
		lca.f[i] = make([]int, lca.m)
	}

	for _, edge := range edges {
		u := edge[0]
		v := edge[1]
		lca.e[u] = append(lca.e[u], v)
		lca.e[v] = append(lca.e[v], u)
	}

	lca.dfs(root, 0)

	for i := 1; i < lca.m; i++ {
		for x := 1; x <= lca.n; x++ {
			lca.f[x][i] = lca.f[lca.f[x][i-1]][i-1]
		}
	}

	return lca
}

func (lca *LCA) dfs(x int, fa int) {
	lca.f[x][0] = fa
	for _, y := range lca.e[x] {
		if y == fa {
			continue
		}
		lca.d[y] = lca.d[x] + 1
		lca.dfs(y, x)
	}
}

func (lca *LCA) Lca(x int, y int) int {
	if lca.d[x] > lca.d[y] {
		x, y = y, x
	}

	for i := lca.m - 1; i >= 0; i-- {
		if lca.d[x] <= lca.d[lca.f[y][i]] {
			y = lca.f[y][i]
		}
	}

	if x == y {
		return x
	}

	for i := lca.m - 1; i >= 0; i-- {
		if lca.f[y][i] != lca.f[x][i] {
			x = lca.f[x][i]
			y = lca.f[y][i]
		}
	}

	return lca.f[x][0]
}

func (lca *LCA) Dis(x int, y int) int {
	return lca.d[x] + lca.d[y] - lca.d[lca.Lca(x, y)]*2
}

const (
	MOD3559 = 1000000007
	N       = 100010
)

var p2 [N]int

func init() {
	p2[0] = 1
	for i := 1; i < N; i++ {
		p2[i] = p2[i-1] * 2 % MOD3559
	}
}

func assignEdgeWeights3559(edges [][]int, queries [][]int) []int {
	lca := NewLCA(edges, 1)
	m := len(queries)
	res := make([]int, m)

	for i := 0; i < m; i++ {
		x := queries[i][0]
		y := queries[i][1]
		if x != y {
			res[i] = p2[lca.Dis(x, y)-1]
		}
	}

	return res
}
