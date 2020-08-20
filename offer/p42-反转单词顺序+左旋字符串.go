package offer

import (
	"strings"
)

// ReverseString reverse a string
//	example:
// 		I am a student.
//		student. a am I
func reverseWord(str string) string {
	strBytes := []byte(str)
	for start, end := 0, len(str)-1; start < end; start, end = start+1, end-1 {
		strBytes[start], strBytes[end] = strBytes[end], strBytes[start]
	}
	return string(strBytes)
}

func ReverseString(str string) string {
	strArray := make([]string, 0)
	for _, word := range strings.Split(reverseWord(str), " ") {
		strArray = append(strArray, reverseWord(word))
	}
	return strings.Join(strArray, " ")
}

// LeftRotateString Rotate k characters to the end.
func LeftRotateString(str string, k int) string {
	// !!!out of range
	strLen := len(str)
	if k <= 0 || k > strLen {
		return str
	}
	reverseString := reverseWord(str)
	return reverseWord(reverseString[:strLen-k]) + reverseWord(reverseString[strLen-k:])
}
