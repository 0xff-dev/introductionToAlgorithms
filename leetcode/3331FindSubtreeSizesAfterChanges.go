package leetcode

func findSubtreeSizes(parent []int, s string) []int {
	n := len(parent)
	if n == 0 {
		return nil
	}

	origAdj := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parent[i]
		origAdj[p] = append(origAdj[p], i)
	}

	last := [26]int{}
	for i := range last {
		last[i] = -1
	}
	newParent := make([]int, n)
	newParent[0] = -1
	var findNewParent func(u int)
	findNewParent = func(u int) {
		c := s[u] - 'a'
		if u > 0 {
			if last[c] != -1 {
				newParent[u] = last[c]
			} else {
				newParent[u] = parent[u]
			}
		}
		prev := last[c]
		last[c] = u
		for _, v := range origAdj[u] {
			findNewParent(v)
		}
		last[c] = prev
	}
	findNewParent(0)
	newAdj := make([][]int, n)
	for i := 1; i < n; i++ {
		p := newParent[i]
		newAdj[p] = append(newAdj[p], i)
	}
	ans := make([]int, n)
	var calcSize func(u int) int
	calcSize = func(u int) int {
		size := 1
		for _, v := range newAdj[u] {
			size += calcSize(v)
		}
		ans[u] = size
		return size
	}

	calcSize(0)
	return ans
}
