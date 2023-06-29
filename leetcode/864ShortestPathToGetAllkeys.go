package leetcode

func shortestPathAllKeys(grid []string) int {
	/*
		x, y := -1, -1
		countOfKey := 0
		rows, cols := len(grid), len(grid[0])
		// 如果是我，我就dfs了直接暴力。
		// 四个方向分别搜一遍，找到最小值，但是估计超时
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				tmp := grid[r][c]
				if tmp == '@' {
					//起点
					x, y = r, c
					continue
				}
				if tmp >= 'a' && tmp <= 'z' {
					countOfKey++
				}
			}
		}
		var (
			dfs  func(int, int, int, int, *int, [][]bool, map[byte]struct{})
			dirs = [][]int{
				{0, 1}, {0, -1}, {1, 0}, {-1, 0},
			}
		)
		dfs = func(x, y, steps, count int, minSteps *int, visited [][]bool, keys map[byte]struct{}) {
			if x < 0 || x >= rows || y < 0 || y >= cols || grid[x][y] == '#' {
				return
			}
			if count == 0 {
				if *minSteps == -1 || *minSteps > steps-1 {
					*minSteps = steps - 1
				}
				return
			}
			if visited[x][y] {
				return
			}

			tmp := grid[x][y]
			del := 0
			if tmp >= 'a' && tmp <= 'z' {
				del = 1
				keys[tmp-'a'] = struct{}{}
			}
			if tmp >= 'A' && tmp <= 'Z' {
				if _, ok := keys[tmp-'A']; !ok {
					return
				}
			}
			visited[x][y] = true

			for _, dir := range dirs {
				nx, ny := x+dir[0], y+dir[1]
				dfs(nx, ny, steps+1, count-del, minSteps, visited, keys)
			}
		}
		ans := -1
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			visited := make([][]bool, rows)
			for r := 0; r < rows; r++ {
				visited[r] = make([]bool, cols)
			}
			visited[x][y] = true
			dfs(nx, ny, 1, countOfKey, &ans, visited, map[byte]struct{}{})
		}
		return ans
	*/
	rows, cols := len(grid), len(grid[0])
	var dirs = [][]int{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
	}
	seen := make(map[int]map[[2]int]struct{})
	keySet := make(map[byte]struct{})
	lockSet := make(map[byte]struct{})
	allKeys := 0
	startR, startC := -1, -1
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			t := grid[r][c]
			if t >= 'a' && t <= 'f' {
				allKeys += (1 << (t - 'a'))
				keySet[t] = struct{}{}
			}
			if t >= 'A' && t <= 'F' {
				lockSet[t] = struct{}{}
			}
			if t == '@' {
				startR, startC = r, c
			}
		}
	}
	queue := [][4]int{{startR, startC, 0, 0}}
	seen[0] = map[[2]int]struct{}{
		[2]int{startR, startC}: struct{}{},
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		curR, curC := cur[0], cur[1]
		keys, dist := cur[2], cur[3]
		for _, dir := range dirs {
			newR, newC := curR+dir[0], curC+dir[1]
			if newR >= 0 && newR < rows && newC >= 0 && newC < cols && grid[newR][newC] != '#' {
				now := grid[newR][newC]
				if _, ok := keySet[now]; ok {
					if (1<<(now-'a'))&keys != 0 {
						continue
					}
					newKeys := (keys | (1 << (now - 'a')))
					if newKeys == allKeys {
						return dist + 1
					}
					if _, ok := seen[newKeys]; !ok {
						seen[newKeys] = make(map[[2]int]struct{})
					}
					seen[newKeys][[2]int{newR, newC}] = struct{}{}
					queue = append(queue, [4]int{newR, newC, newKeys, dist + 1})
				}

				if _, ok := lockSet[now]; ok && (keys&(1<<(now-'A'))) == 0 {
					continue
				} else {
					if v, ok := seen[keys]; ok {
						if _, ok1 := v[[2]int{newR, newC}]; !ok1 {
							seen[keys][[2]int{newR, newC}] = struct{}{}
							queue = append(queue, [4]int{newR, newC, keys, dist + 1})
						}
					}

				}
			}
		}
	}

	return -1
}
