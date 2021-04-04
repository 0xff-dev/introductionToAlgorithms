package leetcode

func spiralOrder(matrix [][]int) []int {
	r := make([]int, 0)
	rows := len(matrix)
	if rows == 0 {
		return r
	}
	cols := len(matrix[0])
	if cols == 0 {
		return r
	}
	rowsLoop := rows / 2
	if rows%2 != 0 {
		rowsLoop++
	}
	colsLoop := cols / 2
	if cols%2 != 0 {
		colsLoop++
	}

	loop := rowsLoop
	if loop > colsLoop {
		loop = colsLoop
	}

	for idx := 0; idx < loop; idx++ {
		x, y := idx, idx
		// right
		for ; y < cols-idx; y++ {
			r = append(r, matrix[x][y])
		}
		x, y = x+1, y-1
		// down
		for ; x < rows-idx; x++ {
			r = append(r, matrix[x][y])
		}
		x, y = x-1, y-1
		if x == idx || y == idx-1 {
			continue
		}
		for ; y >= idx; y-- {
			r = append(r, matrix[x][y])
		}
		x, y = x-1, y+1
		for ; x > idx; x-- {
			r = append(r, matrix[x][y])
		}
	}
	return r
}
