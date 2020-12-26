package datastructure

import (
	"fmt"
)

const INF = 0xffff

type pair struct {
	X, Y int
}

// bfs
// 迷宫最短路径, 从起点开始，向四周发散，当发现可行路径的时候，计算到这里的距离，放入队列。等待下一次的计算。
func bfs(_map [][]byte, dist [][]int, x, y, rows, cols int) int {
	queue := []pair{{x, y}}
	dist[x][y] = 0
	px, py := 0, 0
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		fmt.Println("nowX: ", front.X, " nowY: ", front.Y, " dist: ", dist[front.X][front.Y])
		// if export
		if _map[front.X][front.Y] == 'G' {
			// found
			px, py = front.X, front.Y
			break
		}

		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 || j == 0 {
					nextX, nextY := front.X+i, front.Y+j
					// condition
					if nextX < 0 || nextX >= rows || nextY < 0 || nextY >= cols || _map[nextX][nextY] == '#' || dist[nextX][nextY] != INF {
						continue
					}
					queue = append(queue, pair{nextX, nextY})
					dist[nextX][nextY] = dist[front.X][front.Y] + 1
				}
			}
		}
	}
	return dist[px][py]
}

// 有一组数，是否可以找到和为某个值的几个数字
func dfs(index, sum, k int, nums []int) bool {
	if index == len(nums) {
		return sum == k
	}
	// 下一个选与不选的问题
	if dfs(index+1, sum, k, nums) {
		return true
	}

	if dfs(index+1, sum+nums[index], k, nums) {
		return true
	}

	return false
}

// return bytes subset
func subSetDfs(str []byte) {
	if len(str) == 0 {
		return
	}
	r := make([][]byte, 0)
	for length := 1; length <= len(str); length++ {
		subSet(str, 0, length, []byte{}, &r)
	}
	for idx, sub := range r {
		fmt.Printf("subset[%d]=%s\n", idx, string(sub))
	}
}

func subSet(str []byte, idx, length int, stack []byte, r *[][]byte) {
	if length == 0 {
		_copy := make([]byte, len(stack))
		copy(_copy, stack)
		*r = append(*r, _copy)
		return
	}

	if idx >= len(str) {
		return
	}

	// get
	stack = append(stack, str[idx])
	subSet(str, idx+1, length-1, stack, r)
	stack = stack[:len(stack)-1]
	subSet(str, idx+1, length, stack, r)
}
