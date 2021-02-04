package leetcode

func numTrees(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	array := make([]int, n+1)
	array[0], array[1], array[2], array[3] = 1, 1, 2, 5
	for idx := 4; idx <= n; idx++ {
		tmpSum := 0
		for i := 1; i <= idx; i++ {
			tmpSum += array[i-1] * array[idx-i]
		}
		array[idx] = tmpSum
	}

	return array[n]
}
