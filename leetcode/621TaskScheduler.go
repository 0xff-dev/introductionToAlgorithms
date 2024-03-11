package leetcode

func leastInterval(tasks []byte, n int) int {
	length := len(tasks)

	if n == 0 {
		return length
	}

	taskCount := [26]int{}
	maxFrq := 0
	count := 0
	for _, task := range tasks {
		i := task - 'A'
		taskCount[i]++
		if taskCount[i] == maxFrq {
			count++
		}
		if taskCount[i] > maxFrq {
			maxFrq = taskCount[i]
			count = 1
		}
	}

	ans := (maxFrq-1)*(n+1) + count
	return max(ans, length)
}
