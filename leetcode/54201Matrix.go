package leetcode

func updateMatrix1(mat [][]int) [][]int {
	rows := len(mat)
	cols := len(mat[0])
	ans := make([][]int, rows)
	visited := make([][]bool, rows)
	zeroPoints := make([][2]int, 0)
	for r := 0; r < rows; r++ {
		ans[r] = make([]int, cols)
		visited[r] = make([]bool, cols)
		for c := 0; c < cols; c++ {
			if mat[r][c] == 0 {
				zeroPoints = append(zeroPoints, [2]int{r, c})
			}
		}
	}
	var (
		bfs  func(int, int)
		dirs = [][]int{
			{1, 0}, {0, 1}, {-1, 0}, {0, -1},
		}
	)
	bfs = func(x, y int) {
		distance := 1
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				visited[r][c] = false
			}
		}
		queue := [][2]int{{x, y}}
		visited[x][y] = true
		for len(queue) > 0 {
			next := [][2]int{}
			for _, item := range queue {
				for _, dir := range dirs {
					nx, ny := item[0]+dir[0], item[1]+dir[1]
					if nx < 0 || nx >= rows || ny < 0 || ny >= cols || visited[nx][ny] || mat[nx][ny] == 0 {
						continue
					}
					if mat[nx][ny] == 1 && (ans[nx][ny] == 0 || ans[nx][ny] > distance) {
						ans[nx][ny] = distance
					}
					next = append(next, [2]int{nx, ny})
					visited[nx][ny] = true
				}
			}
			distance++
			queue = next
		}
	}

	for _, point := range zeroPoints {
		x, y := point[0], point[1]
		bfs(x, y)
	}
	/*
		bfs = func(x, y int) int {
			distance := 0
			queue := [][2]int{{x, y}}
			for r := 0; r < rows; r++ {
				for c := 0; c < cols; c++ {
					visited[r][c] = r < x || (r == x && c <= y)
				}
			}
			for len(queue) > 0 {
				next := make([][2]int, 0)
				for _, item := range queue {
					visited[item[0]][item[1]] = true
					if mat[item[0]][item[1]] == 0 {
						return distance
					}
					for _, dir := range dirs {
						nx, ny := item[0]+dir[0], item[1]+dir[1]
						if nx < 0 || nx >= rows || ny < 0 || ny >= cols || visited[nx][ny] {
							continue
						}
						next = append(next, [2]int{nx, ny})
					}
				}
				distance++
				queue = next
			}
			return 0
		}
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				if mat[r][c] == 0 {
					continue
				}
				ans[r][c] = bfs(r, c)
			}
		}
	*/
	return ans
}
