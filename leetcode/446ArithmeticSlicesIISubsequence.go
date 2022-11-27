package leetcode

func numberOfArithmeticSlices(nums []int) int {
	/*
		if len(nums) < 3 {
			return 0
		}
		var dfs func(startIdx, diff, count int, first bool, ans *int)
		dfs = func(startIdx, diff, count int, first bool, ans *int) {
			if startIdx >= len(nums) {
				return
			}
			if count >= 3 {
				*ans++
			}
			for next := startIdx + 1; next < len(nums); next++ {
				if !(first || nums[next]-nums[startIdx] == diff) {
					continue
				}
				diff = nums[next] - nums[startIdx]
				dfs(next, diff, count+1, false, ans)
			}
		}

		ans := 0
		for i := 0; i <= len(nums)-3; i++ {
			dfs(i, 0, 1, true, &ans)
		}
		return ans
	*/
	l := len(nums)
	if l < 3 {
		return 0
	}
	// 到 idx后diff为y有多少个，
	// 1, 2, 3
	diffCache := make(map[int]map[int]int)
	diffCache[l-2] = map[int]int{
		nums[l-1] - nums[l-2]: 1,
	}
	ans := 0

	for idx := l - 3; idx >= 0; idx-- {
		diffCache[idx] = make(map[int]int)

		for next := idx + 1; next < l; next++ {
			diff := nums[next] - nums[idx]
			cal := 1

			if cache, ok := diffCache[next]; ok {
				count, ok := cache[diff]
				if ok {
					cal += count
					ans += cal - 1
				}
			}
			diffCache[idx][diff] += cal
		}
	}
	return ans
}
