package leetcode

func numTeams(rating []int) int {
	lessThan := make([]int, len(rating))
	ans := 0
	for i := 1; i < len(rating); i++ {
		for pre := i - 1; pre >= 0; pre-- {
			if rating[i] > rating[pre] {
				lessThan[i]++
				ans += lessThan[pre]
			}
		}
	}
	lessThan = make([]int, len(rating))
	for i := len(rating) - 2; i >= 0; i-- {
		for next := i + 1; next < len(rating); next++ {
			if rating[i] > rating[next] {
				lessThan[i]++
				ans += lessThan[next]
			}
		}
	}
	return ans
}
