package challengeprogrammingdatastructure

var dirs192 = [][]int{
	{0, 1}, {0, -1},
	{1, 0}, {-1, 0},
}

func toBytes(grid [][]int) []byte {
	bs := make([]byte, 0)
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			bs = append(bs, uint8(grid[row][col])+'0')
		}
	}
	return bs
}

type item struct {
	grid []byte
	zero int
	path int
}

// 9宫格最短路径bfs可以解决掉，如果是16，整体规模超级大。
// 对于16的情况，需要更强劲的剪枝。A*
func Puzzle(grid [][]int) int {
	visited := make(map[string]struct{})
	start := toBytes(grid)
	target := "123456780"
	x, y := 0, 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 0 {
				x, y = row, col
				break
			}
		}
	}

	queue := []item{{start, x*3 + y, 0}}
	visited[string(start)] = struct{}{}
	for len(queue) > 0 {
		next := make([]item, 0)
		for _, x := range queue {
			if string(x.grid) == target {
				return x.path
			}
			for _, dir := range dirs192 {
				x1, y1 := x.zero/3, x.zero%3
				nx, ny := x1+dir[0], y1+dir[1]
				if nx < 0 || nx >= len(grid) || ny <= 0 || ny >= len(grid[0]) {
					continue
				}
				nextBytes := make([]byte, len(x.grid))
				copy(nextBytes, x.grid)
				nextBytes[nx*3+ny], nextBytes[x.zero] = nextBytes[x.zero], nextBytes[nx*3+ny]
				if _, ok := visited[string(nextBytes)]; !ok {
					visited[string(nextBytes)] = struct{}{}
					next = append(next, item{nextBytes, nx*3 + ny, x.path + 1})
				}
			}
		}
		queue = next
	}
	return -1
}
