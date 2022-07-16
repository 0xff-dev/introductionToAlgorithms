package leetcode

const mod576 = 1000000007

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	// timeout
	//var (
	//	findPath func(int, int, int, int, int)
	//	dirs     = [4][2]int{
	//		{-1, 0},
	//		{1, 0},
	//		{0, -1},
	//		{0, 1},
	//	}
	//)
	//ans := 0
	//
	//findPath = func(m, n, maxMove, startRow, startColumn int) {
	//	if startRow >= m || startRow < 0 || startColumn >= n || startColumn < 0 {
	//		ans = (ans + 1) % mod576
	//		// 超出边界
	//		return
	//	}
	//	if maxMove == 0 {
	//		return
	//	}
	//	for _, dir := range dirs {
	//		nextX, nextY := startRow+dir[0], startColumn+dir[1]
	//
	//		findPath(m, n, maxMove-1, nextX, nextY)
	//	}
	//}
	//
	//findPath(m, n, maxMove, startRow, startColumn)
	//return ans

	cache := make([][][]int, m)
	for idx := 0; idx < m; idx++ {
		cache[idx] = make([][]int, n)
		for in := 0; in < n; in++ {
			cache[idx][in] = make([]int, maxMove+1)
			for im := 0; im <= maxMove; im++ {
				cache[idx][in][im] = -1
			}
		}
	}

	var findPath576 func(int, int, int, int, int) int
	findPath576 = func(m int, n int, maxMove int, startRow int, startColumn int) int {
		if startRow == m || startRow < 0 || startColumn == n || startColumn < 0 {
			return 1
		}
		if maxMove == 0 {
			return 0
		}
		if cache[startRow][startColumn][maxMove] == -1 {
			row := (findPath576(m, n, maxMove-1, startRow+1, startColumn) + findPath576(m, n, maxMove-1, startRow-1, startColumn)) % mod576
			col := (findPath576(m, n, maxMove-1, startRow, startColumn+1) + findPath576(m, n, maxMove-1, startRow, startColumn-1)) % mod576
			cache[startRow][startColumn][maxMove] = (row + col) % mod576
		}

		return cache[startRow][startColumn][maxMove]
	}

	return findPath576(m, n, maxMove, startRow, startColumn)
}
