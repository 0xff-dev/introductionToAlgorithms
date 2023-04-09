package challengeprogrammingdatastructure

import (
	"fmt"
	"sort"
)

// 卧槽， 线段树?
// KD Tree
type node142 struct {
	location, p, l, r int
}
type point142 struct {
	id, x, y int
}

func buildKDTree(l, r, depth int, id *int, points []point142, nodes []node142) int {
	if l >= r {
		return -1
	}
	mid := l + (r-l)/2
	if depth&1 == 0 {
		sort.Slice(points[l:r], func(i, j int) bool {
			return points[l+i].x < points[l+j].x
		})
	} else {
		sort.Slice(points[l:r], func(i, j int) bool {
			return points[l+i].y < points[l+j].y
		})
	}
	now := *id
	*id++
	nodes[now].location = mid
	nodes[now].l = buildKDTree(l, mid, depth+1, id, points, nodes)
	nodes[now].r = buildKDTree(mid+1, r, depth+1, id, points, nodes)
	return now
}

func find142(v, sx, tx, sy, ty, depth int, ans *[]point142, points []point142, nodes []node142) {
	x := points[nodes[v].location].x
	y := points[nodes[v].location].y
	if sx <= x && x <= tx && sy <= y && y <= ty {
		*ans = append(*ans, points[nodes[v].location])
	}
	if depth&1 == 0 {
		if nodes[v].l != -1 && sx <= x {
			find142(nodes[v].l, sx, tx, sy, ty, depth+1, ans, points, nodes)
		}
		if nodes[v].r != -1 && x <= tx {
			find142(nodes[v].r, sx, tx, sy, ty, depth+1, ans, points, nodes)
		}
		return
	}
	if nodes[v].l != -1 && sy <= y {
		find142(nodes[v].l, sx, tx, sy, ty, depth+1, ans, points, nodes)
	}
	if nodes[v].r != -1 && ty >= y {
		find142(nodes[v].r, sx, tx, sy, ty, depth+1, ans, points, nodes)
	}
}
func RangeSearch(n int, points1 [][]int, queries [][]int) {
	nodes := make([]node142, n)
	points := make([]point142, n)
	for i := 0; i < n; i++ {
		nodes[i].l, nodes[i].r, nodes[i].p = -1, -1, -1
		points[i] = point142{i, points1[i][0], points1[i][1]}
	}
	id := 0

	root := buildKDTree(0, n, 0, &id, points, nodes)
	for _, q := range queries {
		ans := make([]point142, 0)
		find142(root, q[0], q[1], q[2], q[3], 0, &ans, points, nodes)
		sort.Slice(ans, func(i, j int) bool {
			return ans[i].id < ans[j].id
		})
		for i := 0; i < len(ans); i++ {
			fmt.Printf("id: %d\n", ans[i].id)
		}
		fmt.Println()
	}
}
