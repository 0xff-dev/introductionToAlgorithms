package leetcode

func mincostTickets(days []int, costs []int) int {
	target := make([]int, 366)
	for i := 0; i < 366; i++ {
		target[i] = -1
	}
	dayMap := make(map[int]struct{})
	for _, day := range days {
		dayMap[day] = struct{}{}
	}
	var dp func(int) int

	dp = func(day int) int {
		if day > 365 {
			return 0
		}
		if target[day] != -1 {
			return target[day]
		}
		ans := 0
		_, ok := dayMap[day]
		if ok {
			ans = dp(day+1) + costs[0]
			if r := dp(day+7) + costs[1]; r < ans {
				ans = r
			}
			if r := dp(day+30) + costs[2]; r < ans {
				ans = r
			}
		} else {
			ans = dp(day + 1)
		}
		target[day] = ans
		return ans
	}
	return dp(1)
}
