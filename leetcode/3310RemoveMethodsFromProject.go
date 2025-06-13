package leetcode

type unionFind3310 struct {
	father []int
}

func (u *unionFind3310) findFather(x int) int {
	if x != u.father[x] {
		u.father[x] = u.findFather(u.father[x])
	}
	return u.father[x]
}

func (u *unionFind3310) union(x, y int) {
	fx := u.findFather(x)
	fy := u.findFather(y)
	if fx < fy {
		u.father[fy] = fx
	} else {
		u.father[fx] = fy
	}
}

func remainingMethods(n int, k int, invocations [][]int) []int {
	u := unionFind3310{father: make([]int, n)}
	for i := range n {
		u.father[i] = i
	}
	relations := make(map[int][]int)
	removed := make([]bool, n)
	for _, in := range invocations {
		u.union(in[0], in[1])
		relations[in[0]] = append(relations[in[0]], in[1])
	}
	problemGroup := u.findFather(k)
	removed[k] = true
	var dfs func(int)
	used := map[int]bool{
		k: true,
	}
	dfs = func(cur int) {
		removed[cur] = true
		for _, rel := range relations[cur] {
			if used[rel] {
				continue
			}
			used[rel] = true
			dfs(rel)
		}
	}
	dfs(k)
	var res, keep []int

	for i := range n {
		f := u.findFather(i)
		if f != problemGroup {
			res = append(res, i)
			continue
		}
		if !removed[i] {
			for j := range n {
				keep = append(keep, j)
			}
			return keep
		}
	}
	return res
}
