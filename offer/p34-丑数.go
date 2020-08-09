package offer

import "fmt"

// UglyNum return [index] ugly number
func UglyNum(index int) int {
	if index == 0 {
		return 0
	}
	n2, n3, n5 := 0, 0, 0
	uglyNums := make([]int, index)
	uglyNums[0] = 1
	for cnt := 1; cnt < index; cnt++ {
		num := min(uglyNums[n2]*2, uglyNums[n3]*3, uglyNums[n5]*5)
		uglyNums[cnt] = num
		for uglyNums[n2]*2 <= num {
			n2++
		}
		for uglyNums[n3]*3 <= num {
			n3++
		}
		for uglyNums[n5]*5 <= num {
			n5++
		}
	}
	fmt.Println("ugly num array: ", uglyNums)
	return uglyNums[index-1]
}

func min(args ...int) int {
	if len(args) == 0 {
		return -1
	}
	minNum := args[0]
	for index := 1; index < len(args); index++ {
		if args[index] < minNum {
			minNum = args[index]
		}
	}
	return minNum
}
