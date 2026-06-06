package leetcode

func specialGrid(n int) [][]int {
	if n == 0 {
		return [][]int{{0}}
	}

	size := 1 << n
	ret := make([][]int, size)
	for i := range size {
		ret[i] = make([]int, size)
	}

	var dfs func(int, int, int, int)
	dfs = func(x1, y1, width, start int) {
		if width == 2 {
			ret[x1][y1] = start
			ret[x1+1][y1] = start + 1
			ret[x1+1][y1-1] = start + 2
			ret[x1][y1-1] = start + 3
			return
		}
		mid := width / 2
		nums := mid * mid

		// top right
		dfs(x1, y1, mid, start)
		// top bottom
		dfs(x1+mid, y1, mid, start+nums)
		// left bottom
		dfs(x1+mid, y1-mid, mid, start+nums+nums)
		// left top
		dfs(x1, y1-mid, mid, start+nums+nums+nums)
	}
	dfs(0, size-1, size, 0)
	return ret
}
