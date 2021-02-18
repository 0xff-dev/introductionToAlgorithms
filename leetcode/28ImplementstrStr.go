package leetcode

func nextArr(haystack string) []int {
	nextArray := make([]int, len(haystack))
	nextArray[0] = -1
	compareIndex := -1
	index := 0
	for index < len(haystack)-1 {
		if compareIndex == -1 || haystack[compareIndex] == haystack[index] {
			compareIndex++
			index++
			nextArray[index] = compareIndex
			continue
		}
		compareIndex = nextArray[compareIndex]
	}
	return nextArray
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	haystackLen, needleLen := len(haystack), len(needle)
	nextArray := nextArr(needle)
	i, j := 0, 0
	for i < needleLen && j < haystackLen {
		if i == -1 || needle[i] == haystack[j] {
			i++
			j++
			continue
		}
		i = nextArray[i]
	}
	if i >= needleLen {
		return j - i
	}
	return -1
}
