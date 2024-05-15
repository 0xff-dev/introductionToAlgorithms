package leetcode

import "math"

type pair2812 struct {
	first, second int
}

var dir = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

type solution struct {
	// Directions for moving to neighboring cells: right, left, down, up
	dir [][]int
}

func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)

	multiSourceQueue := make([]pair2812, 0)
	// Mark thieves as 0 and empty cells as -1, and push thieves to the
	// queue
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				// Push thief coordinates to the queue
				multiSourceQueue = append(multiSourceQueue, pair2812{i, j})
				// Mark thief cell with 0
				grid[i][j] = 0
			} else {
				// Mark empty cell with -1
				grid[i][j] = -1
			}
		}
	}

	// Calculate safeness factor for each cell using BFS
	for len(multiSourceQueue) > 0 {
		size := len(multiSourceQueue)
		for size > 0 {
			curr := multiSourceQueue[0]
			multiSourceQueue = multiSourceQueue[1:]
			// Check neighboring cells
			for _, d := range dir {
				di := curr.first + d[0]
				dj := curr.second + d[1]
				val := grid[curr.first][curr.second]
				// Check if the cell is valid and unvisited
				if isValidCell(grid, di, dj) && grid[di][dj] == -1 {
					// Update safeness factor and push to the queue
					grid[di][dj] = val + 1
					multiSourceQueue = append(multiSourceQueue, pair2812{di, dj})
				}
			}
			size--
		}
	}

	// Binary search for maximum safeness factor
	start := 0
	end := 0
	res := -1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// Set end as the maximum safeness factor possible
			end = int(math.Max(float64(end), float64(grid[i][j])))
		}
	}
	for start <= end {
		mid := start + (end-start)/2
		if isValidSafeness(grid, mid) {
			// Store valid safeness and search for larger ones
			res = mid
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return res
}

// Check if a given cell lies within the grid
func isValidCell(grid [][]int, i, j int) bool {
	n := len(grid)
	return i >= 0 && j >= 0 && i < n && j < n
}

// Check if a path exists with given minimum safeness value
func isValidSafeness(grid [][]int, minSafeness int) bool {
	n := len(grid)

	// Check if the source and destination cells satisfy minimum safeness
	if grid[0][0] < minSafeness || grid[n-1][n-1] < minSafeness {
		return false
	}

	traversalQueue := make([]pair2812, 0)
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	traversalQueue = append(traversalQueue, pair2812{0, 0})
	visited[0][0] = true

	// Breadth-first search to find a valid path
	for len(traversalQueue) > 0 {
		curr := traversalQueue[0]
		traversalQueue = traversalQueue[1:]
		if curr.first == n-1 && curr.second == n-1 {
			return true // Valid path found
		}
		// Check neighboring cells
		for _, d := range dir {
			di := curr.first + d[0]
			dj := curr.second + d[1]
			// Check if the neighboring cell is valid and unvisited
			if isValidCell(grid, di, dj) && !visited[di][dj] &&
				grid[di][dj] >= minSafeness {
				visited[di][dj] = true
				traversalQueue = append(traversalQueue, pair2812{di, dj})
			}
		}
	}

	return false // No valid path found
}
