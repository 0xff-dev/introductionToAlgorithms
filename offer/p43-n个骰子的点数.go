package offer

import "fmt"

// recursive
func NDice(n int) {
	result := make(map[int]int)
	auxDice(n, 0, result)
	// calculate percent%%%..
	fmt.Println("result: ", result)
}

func auxDice(n, now int, result map[int]int) {
	if n == 0 {
		result[now]++
		return
	}
	for point := 1; point < 7; point++ {
		auxDice(n-1, now+point, result)
	}
}

// 滚动数组第二次使用
func NDiceLoop(n int) {
	if n < 1 {
		return
	}
	result := make([][]int, 2)
	result[0] = make([]int, 6*n+1)
	result[1] = make([]int, 6*n+1)
	loop := 0
	for index := 1; index <= 6; index++ {
		result[loop][index] = 1
	}
	for index := 2; index <= n; index++ {
		tmpLoop := 1 - loop
		for j := 1; j < index; j++ {
			result[tmpLoop][j] = 0
		}
		for j := index; j <= 6*index; j++ {
			result[tmpLoop][j] = 0
			for inner := 1; inner <= j && inner <= 6; inner++ {
				// 取决与前五个骰子的结果.增加当前的骰子的1-6就是当前的结果
				result[tmpLoop][j] += result[loop][j-inner]
			}
		}
		loop = 1 - loop
	}
	if n%2 == 0 {
		fmt.Println("result: ", result[1])
		return
	}
	fmt.Println("result: ", result[0])
}
