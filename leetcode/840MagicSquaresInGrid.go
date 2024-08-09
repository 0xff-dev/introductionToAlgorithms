package leetcode

func numMagicSquaresInside(grid [][]int) int {
	ans := 0
	for r := 1; r < len(grid)-1; r++ {
		for c := 1; c < len(grid[0])-1; c++ {
			tmp := [10]bool{}
			try := true

			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if !(grid[r+i][c+j] >= 1 && grid[r+i][c+j] <= 9) {
						try = false
						break
					}
					tmp[grid[r+i][c+j]] = true
				}
			}
			if !try {
				continue
			}
			for i := 1; i < 10; i++ {
				if !tmp[i] {
					try = false
					break
				}
			}
			if !try {
				continue
			}
			a := grid[r-1][c-1] + grid[r-1][c] + grid[r-1][c+1]
			b := grid[r+1][c-1] + grid[r+1][c] + grid[r+1][c+1]
			if a != b {
				continue
			}
			b = grid[r][c-1] + grid[r][c] + grid[r][c+1]
			if a != b {
				continue
			}
			// column
			b = grid[r-1][c-1] + grid[r][c-1] + grid[r+1][c-1]
			if a != b {
				continue
			}
			b = grid[r-1][c] + grid[r][c] + grid[r+1][c]
			if a != b {
				continue
			}
			b = grid[r-1][c+1] + grid[r][c+1] + grid[r+1][c+1]
			if a != b {
				continue
			}

			b = grid[r-1][c-1] + grid[r][c] + grid[r+1][c+1]
			if a != b {
				continue
			}
			b = grid[r-1][c+1] + grid[r][c] + grid[r+1][c-1]
			if a != b {
				continue
			}
			ans++
		}
	}
	return ans
}
