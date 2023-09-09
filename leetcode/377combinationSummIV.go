package leetcode

func combinationSum4(nums []int, target int) int {
	var lookup func(int) int
	// 问题是如何做cache
	// 用户 nums 可以组成的元素 1-target 1可以算出来n个，当计算2的时候，可以看取那个，得到她的结果+1
	cache := make(map[int]int)
	lookup = func(sum int) int {
		if sum < 0 {
			return 0
		}
		if sum == 0 {
			return 1
		}
		if v, ok := cache[sum]; ok {
			return v
		}

		ans := 0
		for _, i := range nums {
			next := sum - i
			ans += lookup(next)
		}
		cache[sum] = ans
		return ans
	}
	lookup(target)
	return cache[target]
}
