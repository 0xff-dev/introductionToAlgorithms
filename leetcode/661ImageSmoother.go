package leetcode

func calEege(x, y, rows, cols int) (int, int, int, int, int) {
	tx, ty := x-1, y-1
	bx, by := x+2, y+2
	if tx < 0 {
		tx = 0
	}
	if ty < 0 {
		ty = 0
	}
	if bx > rows {
		bx = rows
	}
	if by > cols {
		by = cols
	}
	count := 9
	if x == 0 && y == 0 || x == 0 && y == cols-1 || x == rows-1 && y == 0 || x == rows-1 && y == cols-1 {
		count = 4
		if rows == 1 || cols == 1 {
			count /= 2
		}
		if rows == 1 && cols == 1 {
			count /= 2
		}
	} else if x == 0 || y == 0 || x == rows-1 || y == cols-1 {
		count = 6
		if rows == 1 || cols == 1 {
			count /= 2
		}
	}
	return tx, ty, bx, by, count
}

func imageSmoother(img [][]int) [][]int {
	rows := len(img)
	cols := len(img[0])
	mat := make([][]int, rows+1)
	for i := 0; i <= rows; i++ {
		mat[i] = make([]int, cols+1)
	}
	for r := rows - 1; r >= 0; r-- {
		for c := cols - 1; c >= 0; c-- {
			mat[r][c] = img[r][c] + mat[r+1][c] + mat[r][c+1] - mat[r+1][c+1]
		}
	}

	result := make([][]int, rows)
	for r := 0; r < rows; r++ {
		result[r] = make([]int, cols)
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			tx, ty, bx, by, count := calEege(r, c, rows, cols)
			result[r][c] = (mat[tx][ty] - mat[tx][by] - mat[bx][ty] + mat[bx][by]) / count
		}
	}
	return result
}
