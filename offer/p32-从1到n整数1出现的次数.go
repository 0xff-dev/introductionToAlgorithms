package offer

import "strconv"

func ExistOneInN(n string) int {
	length := len(n)
	firstNum := n[0] - '0'
	if length == 1 && firstNum == 0 {
		return 0
	}
	if length == 1 && firstNum > 0 {
		return 1
	}
	highNum := 0
	if firstNum > 1 {
		highNum = powerBase(10, len(n)-1)
	} else {
		num, _ := strconv.Atoi(n[1:])
		highNum = num + 1
	}
	lower := int(firstNum) * (length - 1) * powerBase(10, length-2)
	next := ExistOneInN(n[1:])
	return highNum + lower + next
}

func powerBase(base, exp int) int {
	result := 1
	for index := 0; index < exp; index++ {
		result *= base
	}
	return result
}
