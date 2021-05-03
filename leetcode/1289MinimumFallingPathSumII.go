package leetcode

const minSteps = 20000

func minFallingPathSum(arr [][]int) int {
	length := len(arr)
	if length == 0 {
		return 0
	}
	r := [2][]int{}
	for i := 0; i < 2; i++ {
		r[i] = make([]int, length)
		if i == 0 {
			r[i] = arr[0]
			continue
		}
		for j := 0; j < length; j++ {
			r[i][j] = minSteps
		}
	}

	loop := 1
	for row := 1; row < length; row++ {
		// 上一个计算结果, 再使用之后，需要变成minSteps
		for col := 0; col < length; col++ {
			for in := 0; in < length; in++ {
				if in == col {
					continue
				}
				tmp := r[1-loop][col] + arr[row][in]
				if tmp < r[loop][in] {
					r[loop][in] = tmp
				}
			}
			r[1-loop][col] = minSteps
		}
		loop = 1 - loop
	}
	minStep := r[1-loop][0]
	for idx := 1; idx < length; idx++ {
		if r[1-loop][idx] < minStep {
			minStep = r[1-loop][idx]
		}
	}
	return minStep
}
