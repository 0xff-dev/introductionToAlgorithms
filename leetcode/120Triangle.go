package leetcode

func minimumTotal(triangle [][]int) int {
	rows := len(triangle)

	line := make([]int, rows)
	for idx := 0; idx < len(triangle[rows-1]); idx++ {
		line[idx] = triangle[rows-1][idx]
	}

	for row := rows - 2; row >= 0; row-- {
		for col := 0; col <= row; col++ {
			minNum := line[col]
			if minNum > line[col+1] {
				minNum = line[col+1]
			}
			line[col] = triangle[row][col] + minNum
		}
	}
	return line[0]
}

func minimumTotal1(triangle [][]int) int {
	length := len(triangle)
	for row := length - 2; row >= 0; row-- {
		for col := 0; col <= row; col++ {
			_tmp := triangle[row][col] + triangle[row+1][col]
			if r := triangle[row][col] + triangle[row+1][col+1]; r < _tmp {
				_tmp = r
			}
			triangle[row][col] = _tmp
		}
	}
	return triangle[0][0]
}
