package leetcode

func trailingZeroes(n int) int {
	count := 0
	countOfFive := make(map[int]int)
	for step := 5; step <= n; step += 5 {
		countOfFive[step] = 1
		if c, ok := countOfFive[step/5]; ok {
			countOfFive[step] += c
		}
		count += countOfFive[step]
	}
	return count
}
