package offer

import (
	"sort"
)

var pokerMap = map[byte]int{
	'A': 1, '2': 2, '3': 3, '4': 4, '5': 5,
	'6': 6, '7': 7, '8': 8, '9': 9, 'J': 11,
	'Q': 12, 'K': 13, '0': 0,
}

func convertStr2IntArray(str string) []int {
	intArray := make([]int, 0)
	for _, _byte := range []byte(str) {
		if val, ok := pokerMap[_byte]; ok {
			intArray = append(intArray, val)
			continue
		}
		return []int{}
	}
	return intArray
}

func IsStraight(str string) bool {
	if len(str) == 0 || len(str) == 1 {
		return false
	}
	intArray := convertStr2IntArray(str)
	if len(intArray) == 0 {
		return false
	}
	sort.Slice(intArray, func(i, j int) bool {
		return intArray[i] < intArray[j]
	})
	zeroPoker, index, gaps := 0, 0, 0
	for ; index < len(str)-1; index++ {
		if intArray[index] == 0 {
			zeroPoker++
			continue
		}
		if intArray[index] == intArray[index+1] {
			return false
		}
		if intArray[index]+1 != intArray[index+1] {
			gaps += intArray[index+1] - intArray[index] - 1
		}
	}
	if zeroPoker >= gaps {
		return true
	}
	return false
}
