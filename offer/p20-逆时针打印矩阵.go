package offer

import "fmt"

func getLaps(a, b int) int {
	tmpA := a / 2
	if a % 2 != 0 {
		tmpA ++
	}
	tmpB := b / 2
	if b %2 != 0 {
		tmpB ++
	}
	if tmpA > tmpB {
		return tmpB
	}
	return tmpA
}

func CounterclockwisePrint(nums [][]int) {
	// 4个方向
	rows := len(nums)
	if rows == 0 {
		return
	}
	cols := len(nums[0])
	if cols == 0 {
		return
	}
	laps := getLaps(rows, cols)
	for cnt := 0; cnt < laps; cnt++ {
		x, y := cnt, cnt
		printArray(nums, x, y, rows, cols)
	}
}

// 打印外圈数组数字
func printArray(num [][]int, _x, _y, rows, cols int) {
	x, y := _x, _y
	for ; y < cols-_y; y++ {
		fmt.Print(num[x][y])
	}
	x, y = x+1, y-1
	if x >= rows-_x {
		fmt.Println()
		return
	}
	for ; x < rows-_x; x++ {
		fmt.Print(num[x][y])
	}
	x, y = x-1, y-1
	if y < _y {
		fmt.Println()
		return
	}
	for ; y >= _y; y-- {
		fmt.Print(num[x][y])
	}
	x, y = x-1, y+1
	for ; x > _x; x-- {
		fmt.Print(num[x][y])
	}
	fmt.Println()
}
