package leetcode

func numRookCaptures(board [][]byte) int {
	rows, cols := len(board), len(board[0])
	ans := 0
	dirs := [4][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	x, y := -1, -1
	for row := 0; row < rows; row++ {
		col := 0
		for ; col < cols; col++ {
			if board[row][col] == 'R' {
				x, y = row, col
				break
			}
		}
		if col != cols {
			break
		}
	}
	var lookup func(int, int, int, int) int
	lookup = func(x, y, stepX, stepY int) int {
		for !(x >= rows || x < 0 || y >= cols || y < 0 || board[x][y] == 'B') {
			if board[x][y] == 'p' {
				return 1
			}
			x, y = x+stepX, y+stepY
		}
		return 0
	}
	for _, dir := range dirs {
		ans += lookup(x, y, dir[0], dir[1])
	}
	return ans
}
