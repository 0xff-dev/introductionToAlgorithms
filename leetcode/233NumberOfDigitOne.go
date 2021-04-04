package leetcode

import "fmt"

func countDigitOne(n int) int {
	return countDigitOne1([]byte(fmt.Sprintf("%d", n)))
}

func countDigitOne1(bs []byte) int {
	length := len(bs)
	if length == 0 {
		return 0
	}

	highBit := bs[0] - '0'
	if length == 1 {
		if highBit == 0 {
			return 0
		}
		return 1
	}

	high, low := 0, 0
	if highBit == 1 {
		high = getNum(bs[1:]) + 1
	} else if highBit > 1 {
		high = pow10(length - 1)
	}
	low = int(highBit) * pow10(length-2) * (length - 1)

	return high + low + countDigitOne1(bs[1:])
}
func getNum(bs []byte) int {
	base := 0
	for _, item := range bs {
		base = base*10 + int(item-'0')
	}
	return base
}

func pow10(n int) int {
	base := 1
	for c := 1; c <= n; c++ {
		base *= 10
	}
	return base
}
