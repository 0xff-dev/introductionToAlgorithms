package leetcode

func findPoisonedDuration(timeSeries []int, duration int) int {
	ans, length := 0, len(timeSeries)
	if length == 0 {
		return ans
	}
	end := timeSeries[0] + duration - 1
	ans += duration
	// first
	// 1, 2, 3, 4, 5  5
	for idx := 1; idx < length; idx++ {
		if timeSeries[idx] <= end {
			ans -= end - timeSeries[idx] + 1
		}
		end = timeSeries[idx] + duration - 1
		ans += duration
	}
	return ans
}
