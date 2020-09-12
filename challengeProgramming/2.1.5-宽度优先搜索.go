/* 深度优先搜索利用栈，宽度优先搜索利用队列 */

package challengeProgramming

import (
	"fmt"
)

/*
	迷宫由通道和墙壁组成，每次可以向上下左右移动一次，求出从出发点到终点的最短距离。
	假定一定可以从出发点到终点
*/

const INF = 0xffff

type Pair struct {
	X, Y int
}

func bfs(_map [][]byte, dis [][]int, x, y, rows, cols int) int {
	queue := make([]Pair, 0)
	queue = append(queue, Pair{x, y})
	var px, py int
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		if _map[front.X][front.Y] == 'G' {
			px, py = front.X, front.Y
			break
		}
		for row := -1; row <= 1; row++ {
			for col := -1; col <= 1; col++ {
				if row == 0 || col == 0 {
					nextX, nextY := front.X+row, front.Y+col
					// crossover, '#', visited
					if nextX >= 0 && nextX < rows && nextY >= 0 && nextY < cols &&
						_map[nextX][nextY] != '#' && dis[nextX][nextY] == INF {
						queue = append(queue, Pair{nextX, nextY})
						dis[nextX][nextY] = dis[front.X][front.Y] + 1
					}
				}
			}
		}
	}
	return dis[px][py]
}

func MinDistance(_map [][]byte) {
	rows := len(_map)
	if rows == 0 {
		return
	}
	cols := len(_map[0])
	if cols == 0 {
		return
	}
	x, y := 0, 0
	dis := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dis[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dis[i][j] = INF
			if _map[i][j] == 'S' {
				x, y = i, j
			}
		}
	}
	dis[x][y] = 0
	minDis := bfs(_map, dis, x, y, rows, cols)
	if minDis == INF {
		fmt.Println("unreachable")
		return
	}
	fmt.Println("min distance: ", minDis)
}
