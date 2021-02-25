package leetcode

func checkPattern(p string) bool {
	for _, b := range p {
		if b != '*' {
			return false
		}
	}
	return true
}

func isMatch(s string, p string) bool {
	if len(s) == 0 {
		if p == "" || checkPattern(p) {
			return true
		}
		return false
	}
	if len(p) == 0 {
		return false
	}
	rows, cols := len(p)+1, len(s)+1
	r := make([][]uint, rows)
	for idx := 0; idx < rows; idx++ {
		r[idx] = make([]uint, cols)
	}
	r[0][0] = 1
	nextCompareIdx := 1
	for row := 1; row < rows; row++ {
		if nextCompareIdx == cols {
			nextCompareIdx = cols - 1
		}
		if p[row-1] == '*' {
			i := nextCompareIdx
			for ; i < cols && r[row-1][i-1] == 0; i++ {
			}
			if r[row-1][i-1] == 0 {
				break
			}
			for j := i - 1; j < cols; j++ {
				r[row][j] = 1
			}
			continue
		}

		if p[row-1] == '?' {
			for i := nextCompareIdx; i < cols; i++ {
				r[row][i] = 1 & r[row-1][i-1]
			}
			nextCompareIdx++
			continue
		}
		for i := nextCompareIdx; i < cols; i++ {
			if s[i-1] == p[row-1] {
				r[row][i] = 1 & r[row-1][i-1]
			}
		}
		nextCompareIdx++
	}
	return r[rows-1][cols-1] == 1
}
