package leetcode

var dirs = [][2]int{
	{1, 0},
	{0, 1},
	{0, -1},
	{-1, 0},
}

func visitMap(rows, cols int) [][]bool {
	r := make([][]bool, rows)
	for idx := 0; idx < rows; idx++ {
		r[idx] = make([]bool, cols)
	}
	return r
}

func exist(board [][]byte, word string) bool {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			if board[row][col] == word[0] {
				found := false
				visit := visitMap(len(board), len(board[0]))
				visit[row][col] = true
				searchWord(board, word, row, col, 0, visit, &found)
				if found {
					return true
				}
			}
		}
	}
	return false
}

func searchWord(board [][]byte, word string, x, y, idx int, visited [][]bool, r *bool) {
	if *r {
		return
	}
	if idx == len(word) {
		*r = true
		return
	}
	//fmt.Printf("x: %d, y: %d, visit: %v, board: %c, word: %c\n", x, y, visited[x][y], board[x][y], word[idx])
	if word[idx] == board[x][y] {
		visited[x][y] = true
		if idx == len(word)-1 {
			*r = true
			return
		}
		for _, dir := range dirs {
			nextX, nextY := x+dir[0], y+dir[1]
			if nextX < len(board) && nextX >= 0 && nextY < len(board[0]) && nextY >= 0 && !visited[nextX][nextY] {
				searchWord(board, word, nextX, nextY, idx+1, visited, r)
			}
		}
		visited[x][y] = false
	}
}
