package leetcode

type NumMatrix struct {
	matrix [][]int
}

func Constructor304(matrix [][]int) NumMatrix {
	for row := 0; row < len(matrix); row++ {
		for col := 1; col < len(matrix[row]); col++ {
			matrix[row][col] += matrix[row][col-1]
		}
	}

	return NumMatrix{matrix: matrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	ans := 0
	for row := row1; row <= row2; row++ {
		ans += this.matrix[row][col2]
		if col1 > 0 {
			ans -= this.matrix[row][col1-1]
		}
	}
	return ans
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
