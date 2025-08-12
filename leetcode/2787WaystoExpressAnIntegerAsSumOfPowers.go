package leetcode

func power2787(base, n int) int {
	result := base
	for i := 1; i < n; i++ {
		result *= base
	}
	return result
}

const mod2787 = 1000000007

func numberOfWays2787(n int, x int) int {
	candidates := []int{}
	i := 1
	for ; ; i++ {
		p := power2787(i, x)
		if p > n {
			break
		}
		candidates = append(candidates, p)
	}

	var dfs func(index, target int) int
	cache := make(map[[2]int]int)

	dfs = func(index, target int) int {
		if target == 0 {
			return 1
		}
		if target < 0 || index == len(candidates) {
			return 0
		}
		key := [2]int{index, target}
		if val, ok := cache[key]; ok {
			return val
		}
		res := dfs(index+1, target-candidates[index]) % mod2787
		res = (res + dfs(index+1, target)) % mod2787
		cache[key] = res
		return res
	}

	return dfs(0, n)
}
