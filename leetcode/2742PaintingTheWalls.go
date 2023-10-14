package leetcode

const max2742 = 1000000000

func paintWalls(cost []int, time []int) int {

	length := len(cost)
	dp := make([][]int, length+1)
	for i := 0; i <= length; i++ {
		dp[i] = make([]int, length+1)
	}
	for i := 1; i <= length; i++ {
		dp[length][i] = max2742
	}
	for i := length - 1; i >= 0; i-- {
		for remain := 1; remain <= length; remain++ {
			x := 0
			if r := remain - 1 - time[i]; r > x {
				x = r
			}
			paint := cost[i] + dp[i+1][x]
			if dontPaint := dp[i+1][remain]; dontPaint < paint {
				paint = dontPaint
			}
			dp[i][remain] = paint
		}
	}
	return dp[0][length]
}
