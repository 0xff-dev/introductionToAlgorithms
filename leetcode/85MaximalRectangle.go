package leetcode

func maximalRectangle(matrix [][]byte) int {
	ans := 0
	rows, cols := len(matrix), len(matrix[0])
	sum := make([][]int, rows+1)
	for i := 0; i <= rows; i++ {
		sum[i] = make([]int, cols)
	}
	for row := 1; row <= rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row-1][col] == '1' {
				sum[row][col] = sum[row-1][col] + 1
				minHeigh := sum[row][col]
				for pre := col; pre >= 0 && sum[row][pre] != 0; pre-- {
					width := col - pre + 1
					if sum[row][pre] < minHeigh {
						minHeigh = sum[row][pre]
					}
					if r := width * minHeigh; r > ans {
						ans = r
					}
				}
			}
		}
	}
	return ans
}
