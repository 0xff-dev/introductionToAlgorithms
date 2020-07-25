package offer

import "fmt"

// considered n is very large, so we need use unit8 array.
func PrintOne2N(n int) {
	res := make([]uint8, n)
	for isAllNine(res) {
		addOne(res)
		printNum(res)
	}
	fmt.Println("full array")
	FullArray(0, res)
}

func addOne(num []uint8) {
	cf := uint8(1)
	for index := len(num) - 1; index >= 0; index-- {
		num[index] += cf
		cf = 0
		if num[index] >= 10 {
			num[index] -= 10
			cf = 1
		}
	}
}
func isAllNine(num []uint8) bool {
	for _, _uint8Num := range num {
		if _uint8Num != 9 {
			return true
		}
	}
	return false
}

func printNum(num []uint8) {
	first := true
	for index := 0; index < len(num); index++ {
		if num[index] == 0 && first {
			continue
		}
		first = false
		fmt.Printf("%d", num[index])
	}
	fmt.Println()
}

// 全排列解法
func FullArray(start int, num []uint8) {
	if start == len(num) {
		printNum(num)
		return
	}
	var n uint8 = 0
	for ; n < 10; n++ {
		num[start] = n
		FullArray(start+1, num)
	}
}
