// n 皇后问题
package algorith

import "fmt"

func NEmpress(n int) {
	cnt := 0
	arr := make([]int, n)
	aux(0, n, &cnt, arr)
	fmt.Println("all cnt", cnt)
}

func aux(row, n int, cnt *int, arr []int) {
	if row == n {
		fmt.Println(arr)
		for _, item := range arr {
			for index := 0; index < n; index++ {
				if index == item {
					fmt.Print("Q")
					continue
				}
				fmt.Print(".")
			}
			fmt.Println()
		}

		*cnt++
		return
	}
	for index := 0; index < n; index++ {
		arr[row] = index // 第row，可以放置的位置
		ok := true
		for _row := 0; _row < row; _row++ {
			// 同行，同斜行判断
			if arr[_row] == arr[row] || arr[row]+_row == arr[_row]+row || arr[_row]-row == arr[row]-_row {
				ok = false
				break
			}
		}
		if ok {
			aux(row+1, n, cnt, arr)
		}
	}
}
