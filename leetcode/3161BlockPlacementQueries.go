package leetcode

import "github.com/emirpasic/gods/trees/redblacktree"

func getResults3161(queries [][]int) []bool {
	const mx = 50000
	seg := make([]int, mx<<2)

	var update func(idx, val, p, l, r int)
	update = func(idx, val, p, l, r int) {
		if l == r {
			seg[p] = val
			return
		}
		mid := (l + r) >> 1
		if idx <= mid {
			update(idx, val, p<<1, l, mid)
		} else {
			update(idx, val, p<<1|1, mid+1, r)
		}
		if seg[p<<1] > seg[p<<1|1] {
			seg[p] = seg[p<<1]
		} else {
			seg[p] = seg[p<<1|1]
		}
	}

	var query func(L, R, p, l, r int) int
	query = func(L, R, p, l, r int) int {
		if L <= l && r <= R {
			return seg[p]
		}
		mid := (l + r) >> 1
		res := 0
		if L <= mid {
			if val := query(L, R, p<<1, l, mid); val > res {
				res = val
			}
		}
		if R > mid {
			if val := query(L, R, p<<1|1, mid+1, r); val > res {
				res = val
			}
		}
		return res
	}

	tree := redblacktree.NewWithIntComparator()
	tree.Put(0, struct{}{})
	tree.Put(mx, struct{}{})
	update(mx, mx, 1, 0, mx)
	ans := make([]bool, 0, len(queries))

	for _, q := range queries {
		if q[0] == 1 {
			x := q[1]
			r := mx
			if node, ok := tree.Ceiling(x + 1); ok {
				r = node.Key.(int)
			}
			l := 0
			if node, ok := tree.Floor(x - 1); ok {
				l = node.Key.(int)
			}

			update(x, x-l, 1, 0, mx)
			update(r, r-x, 1, 0, mx)
			tree.Put(x, struct{}{})
		} else {
			x, sz := q[1], q[2]
			pre := 0
			if node, ok := tree.Floor(x); ok {
				pre = node.Key.(int)
			}
			maxSpace := query(0, pre, 1, 0, mx)
			if x-pre > maxSpace {
				maxSpace = x - pre
			}
			ans = append(ans, maxSpace >= sz)
		}
	}

	return ans
}
