package offer

import "fmt"

// 全排列
func StrAllSort(str string) {
	byteAux([]byte(str), 0)
}

func byteAux(strBytes []byte, location int) {
	if location == len(strBytes) {
		fmt.Println("an order: ", string(strBytes))
		return
	}
	for index := location; index < len(strBytes); index++ {
		strBytes[index], strBytes[location] = strBytes[location], strBytes[index]
		byteAux(strBytes, location+1)
		strBytes[index], strBytes[location] = strBytes[location], strBytes[index]
	}
}

// 字符串子集
func StrSubSet(str string) {
	// 子集长度1-len(str)-1
	for subSetLen := 1; subSetLen <= len(str); subSetLen++ {
		subSetAux(0, subSetLen, str, []byte{})
	}
}

func subSetAux(start, length int, str string, stack []byte) {
	if length == 0 {
		fmt.Println("subset is: ", string(stack))
		return
	}
	if start >= len(str) {
		return
	}
	stack = append(stack, str[start])
	subSetAux(start+1, length-1, str, stack)
	stack = stack[:len(stack)-1]
	subSetAux(start+1, length, str, stack)
}

// 八皇后问题
// 同直线，同斜线不可有两个皇后，问有几种方法排列皇后
func NEmpress(n int) {
	solution := 0
	location := make([]int, n)
	empressAux(0, n, location, &solution)
	fmt.Println("solution: ", solution)
}

func empressAux(row, n int, location []int, solution *int) {
	if row == n {
		*solution++
		empressLocationPrint(location, n)
		return
	}
	for col := 0; col < n; col++ {
		location[row] = col // 在row层，皇后放到哪个位置
		ok := true
		for _row := 0; _row < row; _row++ {
			if location[_row] == location[row] || location[row]+row == location[_row]+_row || location[row]-row == location[_row]-_row {
				ok = false
				break
			}
		}
		if ok {
			empressAux(row+1, n, location, solution)
		}
	}
}

func empressLocationPrint(location []int, n int) {
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if location[row] == col {
				fmt.Print("Q")
				continue
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
	fmt.Println()
}
