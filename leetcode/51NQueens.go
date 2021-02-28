package leetcode

import "bytes"

func solveNQueens(n int) [][]string {
	r := make([][]string, 0)
	path := make([]int, n)
	nQueens(0, n, path, &r)
	return r
}

func nQueens(row, n int, path []int, r *[][]string) {
	if row == n {
		// success queens setting.
		line := make([]string, 0)
		for _, queensLocation := range path {
			buf := bytes.NewBuffer([]byte{})
			for idx := 0; idx < n; idx++ {
				if idx != queensLocation {
					buf.WriteByte('.')
					continue
				}
				buf.WriteByte('Q')
			}
			line = append(line, buf.String())
		}
		*r = append(*r, line)
		return
	}

	for col := 0; col < n; col++ {
		path[row] = col
		ok := true
		// if condition {}
		for _row := 0; _row < row; _row++ {
			// 同行，同斜行判断
			if path[_row] == path[row] || path[row]+_row == path[_row]+row || path[_row]-row == path[row]-_row {
				ok = false
				break
			}
		}
		if ok {
			nQueens(row+1, n, path, r)
		}
	}
}
