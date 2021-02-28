package leetcode

func totalNQueens(n int) int {
	r := 0
	path := make([]int, n)
	nQueensCount(0, n, path, &r)
	return r
}

func nQueensCount(row, n int, path []int, r *int) {
	if row == n {
		*r++
		return
	}
	for col := 0; col < n; col++ {
		path[row] = col
		ok := true
		for _row := 0; _row < row; _row++ {
			if path[_row] == path[row] || path[_row]+row == path[row]+_row || path[row]-_row == path[_row]-row {
				ok = false
				break
			}
		}
		if ok {
			nQueensCount(row+1, n, path, r)
		}
	}
}
