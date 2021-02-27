package leetcode

// n*n
func rotate(matrix [][]int) {
	rotateCount := len(matrix)
	n := len(matrix)
	idx := 0
	for rotateCount > 1 {
		// first element.
		count := rotateCount - 1
		for count > 0 {
			x, y, now := idx+1, idx, matrix[idx][idx]
			for ; x < n-idx; x++ {
				matrix[x-1][y] = matrix[x][y]
			}
			x, y = x-1, y+1
			for ; y < n-idx; y++ {
				matrix[x][y-1] = matrix[x][y]
			}
			y, x = y-1, x-1
			for ; x >= idx; x-- {
				matrix[x+1][y] = matrix[x][y]
			}
			x, y = x+1, y-1
			for ; y > idx; y-- {
				matrix[x][y+1] = matrix[x][y]
			}
			matrix[idx][y+1] = now
			count--
		}
		rotateCount -= 2
		idx++
	}
}
