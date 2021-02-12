package leetcode

import (
	"fmt"
)

func compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}
	nowByteIndex, sameCount := 0, 1
	for index := 1; index < len(chars); {
		if chars[index] == chars[nowByteIndex] {
			sameCount++
			index++
			continue
		}
		if sameCount == 1 {
			nowByteIndex = index
			index++
			continue
		}
		remaining := chars[index:]
		count := []byte(fmt.Sprintf("%d", sameCount))
		chars = append(chars[:nowByteIndex+1], count...)
		chars = append(chars, remaining...)
		nowByteIndex = nowByteIndex + len(count) + 1
		index = nowByteIndex + 1
		sameCount = 1
	}
	if sameCount != 1 {
		count := []byte(fmt.Sprintf("%d", sameCount))
		chars = append(chars[:nowByteIndex+1], count...)
	}
	return len(chars)
}

//func compress(chars []byte) int {
//	if len(chars) == 0 {
//		return 0
//	}
//
//	length, nowByte, sameCount := 0, chars[0], 1
//
//	for index := 1; index < len(chars); index++ {
//		if chars[index] == nowByte {
//			sameCount++
//			continue
//		}
//		nowByte = chars[index]
//		if sameCount == 1 {
//			length++
//			continue
//		}
//		length += len(fmt.Sprintf("%d", sameCount)) + 1
//		sameCount = 1
//	}
//	if sameCount == 1 {
//		length++
//	} else {
//		length += len(fmt.Sprintf("%d", sameCount)) + 1
//	}
//
//	return length
//}
