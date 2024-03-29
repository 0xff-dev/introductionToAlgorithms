package leetcode

func numPairsDivisibleBy60(time []int) int {
	durations := make(map[int]int)
	durations[time[0]] = 1
	ans := 0
	nums := make([]int, 0)
	for i := 0; i <= 1000; i += 60 {
		nums = append(nums, i)
	}
	for i := 1; i < len(time); i++ {
		j := len(nums) - 1
		for ; j >= 0 && nums[j] >= time[i]; j-- {
			if c, ok := durations[nums[j]-time[i]]; ok {
				ans += c
			}
		}
		durations[time[i]]++
	}
	return ans
}

func numPairsDivisibleBy602(time []int) int {
	cache := [60]int{}
	ans := 0
	for _, dur := range time {
		v := dur % 60
		if v == 0 {
			ans += cache[0]
		} else {
			ans += cache[60-v]
		}
		cache[v]++
	}
	return ans
}
