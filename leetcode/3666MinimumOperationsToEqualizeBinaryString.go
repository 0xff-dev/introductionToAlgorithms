package leetcode

import (
	"math"

	"github.com/emirpasic/gods/trees/redblacktree"
)

func minOperations3666(s string, k int) int {
	n, m := len(s), 0
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	nodeSets := make([]*redblacktree.Tree, 2)
	nodeSets[0] = redblacktree.NewWithIntComparator()
	nodeSets[1] = redblacktree.NewWithIntComparator()
	for i := 0; i <= n; i++ {
		nodeSets[i%2].Put(i, struct{}{})
		if i < n && s[i] == '0' {
			m++
		}
	}
	q := []int{m}
	dist[m] = 0
	nodeSets[m%2].Remove(m)
	for len(q) > 0 {
		m := q[0]
		q = q[1:]
		c1, c2 := max(k-n+m, 0), min(k, m)
		lnode, rnode := m+k-2*c2, m+k-2*c1
		nodeSet := nodeSets[lnode%2]
		for m2, found := nodeSet.Ceiling(lnode); found && m2.Key.(int) <= rnode; m2, found = nodeSet.Ceiling(lnode) {
			dist[m2.Key.(int)] = dist[m] + 1
			q = append(q, m2.Key.(int))
			nodeSet.Remove(m2.Key.(int))
		}
	}
	if dist[0] == math.MaxInt32 {
		return -1
	}
	return dist[0]
}
