package leetcode

func matrixReshape(mat [][]int, r int, c int) [][]int {
	rows, cols := len(mat), len(mat[0])
	if r*c != rows*cols {
		return mat
	}
	res := make([][]int, r)
	for i := range r {
		res[i] = make([]int, c)
	}
	x, y := 0, 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			res[x][y] = mat[i][j]
			y++
			if y == c {
				x++
				y = 0
			}
		}
	}
	return res
}
