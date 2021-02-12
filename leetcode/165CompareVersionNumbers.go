package leetcode

import (
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	intV1, intV2 := checkVersion(v1), checkVersion(v2)
	lenV1, lenV2 := len(intV1), len(intV2)
	if lenV1 > lenV2 {
		for count := 0; count < lenV1-lenV2; count++ {
			intV2 = append(intV2, 0)
		}
	} else {
		for count := 0; count < lenV2-lenV1; count++ {
			intV1 = append(intV1, 0)
		}
	}
	for idx := 0; idx < len(intV1); idx++ {
		if intV1[idx] > intV2[idx] {
			return 1
		}
		if intV1[idx] < intV2[idx] {
			return -1
		}
	}
	return 0
}

func checkVersion(version []string) []int {
	r := make([]int, 0)
	for idx := 0; idx < len(version); idx++ {
		charIndex := 0
		for ; charIndex < len(version[idx]) && version[idx][charIndex] == '0'; charIndex++ {
		}
		if charIndex == len(version[idx]) {
			r = append(r, 0)
			continue
		}

		r = append(r, toInt(version[idx][charIndex:]))
	}
	return r
}

func toInt(str string) int {
	base := 0
	for _, b := range str {
		base = base*10 + int(b-'0')
	}
	return base
}
