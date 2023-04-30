package leetcode

type unionFind1579 struct {
	representative []int
	componentSize  []int
	components     int
}

func (u *unionFind1579) findRepresentative(x int) int {
	if u.representative[x] != x {
		u.representative[x] = u.findRepresentative(u.representative[x])
	}
	return u.representative[x]
}

func (u *unionFind1579) union(x, y int) int {
	fx := u.findRepresentative(x)
	fy := u.findRepresentative(y)
	if fx == fy {
		return 0
	}
	if u.componentSize[fx] > u.componentSize[fy] {
		u.componentSize[fx] += u.componentSize[fy]
		u.representative[fy] = fx
	} else {
		u.componentSize[fy] += u.componentSize[fx]
		u.representative[fx] = fy
	}
	u.components--
	return 1
}

func (u *unionFind1579) isConnected() bool {
	return u.components == 1
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
	a := unionFind1579{
		representative: make([]int, n+1),
		componentSize:  make([]int, n+1),
		components:     n,
	}
	b := unionFind1579{
		representative: make([]int, n+1),
		componentSize:  make([]int, n+1),
		components:     n,
	}
	for i := 0; i <= n; i++ {
		a.representative[i] = i
		b.representative[i] = i
		a.componentSize[i] = 1
		b.componentSize[i] = 1
	}

	edgeRequired := 0
	for _, e := range edges {
		whichType, from, to := e[0], e[1], e[2]
		if whichType == 3 {
			// both
			edgeRequired += (a.union(from, to) | b.union(from, to))
		}
	}
	for _, e := range edges {
		whichType, from, to := e[0], e[1], e[2]
		if whichType == 1 {
			edgeRequired += a.union(from, to)
			continue
		}
		edgeRequired += b.union(from, to)
	}
	if a.isConnected() && b.isConnected() {
		return len(edges) - edgeRequired
	}
	return -1
}
