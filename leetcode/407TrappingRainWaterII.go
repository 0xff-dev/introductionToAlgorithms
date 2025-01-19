package leetcode

import (
	"container/heap"
)

type Cell struct {
	height int
	row    int
	col    int
}

// A min-heap of Cells
type MinHeap407 []Cell

func (h MinHeap407) Len() int           { return len(h) }
func (h MinHeap407) Less(i, j int) bool { return h[i].height < h[j].height }
func (h MinHeap407) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap407) Push(x interface{}) {
	*h = append(*h, x.(Cell))
}

func (h *MinHeap407) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func trapRainWater(heightMap [][]int) int {
	if len(heightMap) == 0 || len(heightMap[0]) == 0 {
		return 0
	}

	dRow := []int{0, 0, -1, 1}
	dCol := []int{-1, 1, 0, 0}
	numOfRows := len(heightMap)
	numOfCols := len(heightMap[0])

	visited := make([][]bool, numOfRows)
	for i := range visited {
		visited[i] = make([]bool, numOfCols)
	}

	boundary := &MinHeap407{}
	heap.Init(boundary)

	// Add the first and last column cells to the boundary
	for i := 0; i < numOfRows; i++ {
		heap.Push(boundary, Cell{height: heightMap[i][0], row: i, col: 0})
		heap.Push(boundary, Cell{height: heightMap[i][numOfCols-1], row: i, col: numOfCols - 1})
		visited[i][0] = true
		visited[i][numOfCols-1] = true
	}

	// Add the first and last row cells to the boundary
	for i := 0; i < numOfCols; i++ {
		heap.Push(boundary, Cell{height: heightMap[0][i], row: 0, col: i})
		heap.Push(boundary, Cell{height: heightMap[numOfRows-1][i], row: numOfRows - 1, col: i})
		visited[0][i] = true
		visited[numOfRows-1][i] = true
	}

	totalWaterVolume := 0

	for boundary.Len() > 0 {
		currentCell := heap.Pop(boundary).(Cell)
		currentRow, currentCol, minBoundaryHeight := currentCell.row, currentCell.col, currentCell.height

		for direction := 0; direction < 4; direction++ {
			neighborRow := currentRow + dRow[direction]
			neighborCol := currentCol + dCol[direction]

			if isValidCell407(neighborRow, neighborCol, numOfRows, numOfCols) && !visited[neighborRow][neighborCol] {
				neighborHeight := heightMap[neighborRow][neighborCol]

				// If the neighbor's height is less than the current boundary height, water can be trapped
				if neighborHeight < minBoundaryHeight {
					totalWaterVolume += minBoundaryHeight - neighborHeight
				}

				// Push the neighbor into the boundary with updated height
				heap.Push(boundary, Cell{
					height: max(neighborHeight, minBoundaryHeight),
					row:    neighborRow,
					col:    neighborCol,
				})
				visited[neighborRow][neighborCol] = true
			}
		}
	}

	return totalWaterVolume
}

func isValidCell407(row, col, numOfRows, numOfCols int) bool {
	return row >= 0 && col >= 0 && row < numOfRows && col < numOfCols
}
