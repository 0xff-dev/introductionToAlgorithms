package leetcode

/*
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
*/

type NumMatrix struct {
	m [][]int
}

func Constructor304(matrix [][]int) NumMatrix {
	/*

	  1 2
	  3 4

	*/
	rows, cols := len(matrix), len(matrix[0])
	m := make([][]int, rows+1)
	for i := 0; i <= len(matrix); i++ {
		m[i] = make([]int, cols+1)
	}
	for row := rows - 1; row >= 0; row-- {
		for col := cols - 1; col >= 0; col-- {
			m[row][col] = matrix[row][col] + m[row+1][col] + m[row][col+1] - m[row+1][col+1]
		}
	}
	return NumMatrix{m: m}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.m[row1][col1] - this.m[row2+1][col1] - this.m[row1][col2+1] + this.m[row2+1][col2+1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
