package leetcode

func generate(numRows int) [][]int {
	r := make([][]int, 0)
	if numRows < 1 {
		return r
	}

	r = append(r, []int{1})
	for row := 2; row <= numRows; row++ {
		tmp := make([]int, row)
		tmp[0], tmp[row-1] = 1, 1
		for idx := 1; idx < row-1; idx++ {
			tmp[idx] = r[row-2][idx-1] + r[row-2][idx]
		}
		r = append(r, tmp)
	}

	return r
}
