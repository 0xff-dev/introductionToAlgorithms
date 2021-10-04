package leetcode

var localDirs = [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func floodFill(image [][]int, sr, sc int, newColor int) [][]int {
	rows, cols := len(image), len(image[0])
	visited := make([][]int, len(image))
	for idx := 0; idx < rows; idx++ {
		visited[idx] = make([]int, cols)
	}

	floodFillDFS(image, visited, sr, sc, rows, cols, image[sr][sc], newColor)
	return image
}

// 广度优先遍历
func floodFillDFS(image, visited [][]int, sr, sc, rows, cols int, sourceColor, newColor int) {
	if sr < 0 || sr >= rows || sc < 0 || sc >= cols {
		return
	}

	if visited[sr][sc] > 0 || image[sr][sc] != sourceColor {
		return
	}

	image[sr][sc] = newColor
	visited[sr][sc] = 1
	for _, dir := range localDirs {
		floodFillDFS(image, visited, sr+dir[0], sc+dir[1], rows, cols, sourceColor, newColor)
	}
	// for r := -1; r <= 1; r++ {
	// 	for c := -1; c <= 1; c++ {
	// 		if r != 0 || c != 0 {
	// 			floodFillDFS(image, visited, sr+r, sc+c, rows, cols, sourceColor, newColor)
	// 		}
	// 	}
	// }
}
