/*
	一组数据，从中选取若干，让他们的和为k, 深度搜索
*/
package challengeProgramming

import "fmt"

//1 深度搜索
func FindSomeK(k int, nums []int) {
	dfs(0, k, 0, nums)
}

func dfs(start, k, nowSum int, nums []int) bool {
	if start == len(nums) {
		return nowSum == k
	}
	if dfs(start+1, k, nowSum, nums) {
		return true
	}
	if dfs(start+1, k, nowSum+nums[start], nums) {
		return true
	}
	return false
}

func LakeCount(puddle [][]byte) {
	if len(puddle) == 0 {
		return
	}
	if len(puddle[0]) == 0 {
		return
	}
	count := 0
	row, col := len(puddle), len(puddle[0])
	for x := 0; x < row; x++ {
		for y := 0; y < col; y++ {
			if puddle[x][y] == 'w' {
				lakeDfs(x, y, row, col, puddle)
				count++
			}
		}
	}
	fmt.Println("result: ", count)
}

func lakeDfs(x, y, row, col int, puddle [][]byte) {
	puddle[x][y] = '.'
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			nextX, nextY := x+i, y+j
			if nextX >= 0 && nextX < row && nextY >= 0 && nextY < col && puddle[nextX][nextY] == 'w' {
				lakeDfs(nextX, nextY, row, col, puddle)
			}
		}
	}
}
