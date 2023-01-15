package leetcode

func countNegatives(grid [][]int) int {
	ans := 0
	for row := 0; row < len(grid); row++ {
		// 4, 3, 2, -1
		l, i := len(grid[row]), 0
		for ; i < l; i++ {
			if grid[row][i] < 0 {
				break
			}
		}
		ans += l - i
	}
	return ans
}
