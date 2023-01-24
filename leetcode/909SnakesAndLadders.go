package leetcode

// type cellPath struct {
// 	obj            int
// 	path           int
// 	snakerOrLadder bool
// }

// type cellPathList []cellPath

// func (cl *cellPathList) Len() int {
// 	return len(*cl)
// }

// func (cl *cellPathList) Less(i, j int) bool {
// 	if (*cl)[i].path == (*cl)[j].path {
// 		return (*cl)[i].obj > (*cl)[j].obj
// 	}
// 	return (*cl)[i].path < (*cl)[j].path
// }

// func (cl *cellPathList) Swap(i, j int) {
// 	(*cl)[i], (*cl)[j] = (*cl)[j], (*cl)[i]
// }

// func (cl *cellPathList) Push(x interface{}) {
// 	*cl = append(*cl, x.(cellPath))
// }

// func (cl *cellPathList) Pop() interface{} {
// 	old := *cl
// 	l := len(old)
// 	x := old[l-1]
// 	*cl = old[:l-1]
// 	return x
// }

func snakesAndLadders(board [][]int) int {
	n := len(board)
	target := n * n
	var position func(int) (int, int)
	position = func(x int) (int, int) {
		divisor := x / n
		quotient := x % n
		row := n - divisor - 1
		col := n - 1
		if quotient == 0 {
			divisor--
			row++
			if divisor&1 == 1 {
				col = 0
			}
			return row, col
		}
		col = quotient - 1
		if divisor&1 == 1 {
			col = n - 1 - col
		}
		return row, col
	}

	/*
		queue := cellPathList([]cellPath{{1, 0, false}})
		heap.Init(&queue)

		ans := -1
		visited := make([]bool, n*n+1)
		for len(queue) > 0 {
			item := heap.Pop(&queue).(cellPath)
			if visited[item.obj] {
				continue
			}
			visited[item.obj] = true
			if item.obj == target {
				if ans == -1 || ans > item.path {
					ans = item.path
				}
				continue
			}

			x, y := position(item.obj)
			snakerOrLadder := board[x][y] != -1
			if snakerOrLadder && !item.snakerOrLadder {
				nx, ny := position(board[x][y])
				skip := false
				if board[nx][ny] != -1 {
					skip = true
				}
				heap.Push(&queue, cellPath{board[x][y], item.path, skip})
				continue
			}
			/*
				if snakerOrLadder && !item.snakerOrLadder {
					heap.Push(&queue, cellPath{board[x][y], item.path, true})
					continue
				}

			for next := item.obj + 1; next <= item.obj+6 && next <= target; next++ {
				if next == board[x][y] && item.snakerOrLadder {
					continue
				}
				heap.Push(&queue, cellPath{next, item.path + 1, false})
			}
		}
		return ans
	*/

	queue := []int{1}
	dist := make([]int, target+1)
	for i := 0; i <= target; i++ {
		dist[i] = -1
	}
	dist[1] = 0
	for len(queue) > 0 {
		x := queue[0]
		queue = queue[1:]
		for next := x + 1; next <= x+6 && next <= target; next++ {
			nx, ny := position(next)
			destination := next
			if board[nx][ny] != -1 {
				destination = board[nx][ny]
			}

			if dist[destination] == -1 {
				dist[destination] = dist[x] + 1
				queue = append(queue, destination)
			}
		}
	}
	return dist[target]
}
