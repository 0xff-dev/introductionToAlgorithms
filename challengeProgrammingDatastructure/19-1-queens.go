package challengeprogrammingdatastructure

import "fmt"

func Queens(n int) int {
	pos := make([]int, n)
	ans := 0
	queenDFS(n, 0, pos, &ans)
	return ans
}

func queenDFS(n, now int, pos []int, ans *int) {
	if now >= n {
		*ans++
		fmt.Printf("---- %d\n", *ans)
		PrintAnswer(pos, n)
		return
	}
	for col := 0; col < n; col++ {
		// 设置哪个位置
		pos[now] = col
		ok := true
		for pre := 0; pre < now; pre++ {
			cond1 := col == pos[pre]
			cond2 := now-pre == col-pos[pre]
			cond3 := now-pre+col-pos[pre] == 0
			if cond1 || cond2 || cond3 {
				ok = false
				break
			}
		}
		if ok {
			queenDFS(n, now+1, pos, ans)
		}
	}
}
func PrintAnswer(pos []int, n int) {
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if pos[row] == col {
				fmt.Print("Q ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
