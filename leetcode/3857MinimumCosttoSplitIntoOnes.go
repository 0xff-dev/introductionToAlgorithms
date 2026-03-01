package leetcode

func minCost3857(n int) int {
	cache := map[int]int{
		1: 0,
	}
	var dfs func(int) int
	dfs = func(x int) int {
		if v, ok := cache[x]; ok {
			return v
		}
		ret := (1 << 31) - 1
		cost := 0
		for i := 1; i <= x/2; i++ {
			cost = i * (x - i)
			ret = min(ret, dfs(i)+dfs(x-i)+cost)
		}
		cache[x] = ret
		return ret
	}
	return dfs(n)
}
