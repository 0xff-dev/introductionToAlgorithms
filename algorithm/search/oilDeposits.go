package search

import "fmt"

/*
	输入一个m*n的字符矩阵，统计@组成多少个连同八块，两个各自相邻(上下/对角线)属于同一个连同八块
*/
const (
	maxc = 1000
	maxr = 1000
)

var idx [maxr][maxc]int
var char [maxr][maxc]byte // store byte
var m, n int

func dfs(r, c, cnt int) {
	if r < 0 || r >= m || c < 0 || c >= n {
		return // 边界条件
	}
	if char[r][c] != '@' || idx[r][c] > 0 {
		return // 被访问过或者是非@字符，不予以处理
	}
	idx[r][c] = cnt
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dc != 0 || dr != 0 {
				dfs(r+dr, c+dc, cnt)
			}
		}
	}
}

func oil() {
	for n, err := fmt.Scanf("%d %d", &m, &n); n == 2 && err == nil; {
		for r := 0; r < m; r++ {
			for c := 0; c < n; c++ {
				_, _ = fmt.Scanf("%c", &char[r][c])
			}
		}
		cnt := 0
		for r := 0; r < m; r++ {
			for c := 0; c < n; c++ {
				if char[r][c] == '@' && idx[r][c] == 0 {
					cnt++
					dfs(r, c, cnt)
				}
			}
		}
		fmt.Println("count: ", cnt)
	}
}
