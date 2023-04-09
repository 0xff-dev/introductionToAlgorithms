package challengeprogrammingdatastructure

type unionFind141 struct {
	father []int
}

func (u *unionFind141) findFather(x int) int {
	if u.father[x] != x {
		u.father[x] = u.findFather(u.father[x])
	}
	return u.father[x]
}

func (u *unionFind141) union(x, y int) {
	fx := u.findFather(x)
	fy := u.findFather(y)
	if fx < fy {
		u.father[y] = fx
	} else {
		u.father[x] = fy
	}
}

func UnionFindTree(n, q int, options [][]int) []int {
	u := unionFind141{father: make([]int, n)}
	for i := 0; i < n; i++ {
		u.father[i] = i
	}
	ans := make([]int, 0)
	for _, op := range options {
		a, b := op[1], op[2]
		fa := u.findFather(a)
		fb := u.findFather(b)
		if op[0] == 0 {
			u.union(fa, fb)
			continue
		}
		same := 1
		if fa != fb {
			same = 0
		}
		ans = append(ans, same)
	}
	return ans
}
