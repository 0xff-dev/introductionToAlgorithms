package leetcode

func gcdOfStrings(str1 string, str2 string) string {
	ls1, ls2 := len(str1), len(str2)
	// 找到你来的所有公约数
	common := make([]int, 0)
	for i := 1; i <= ls1 && i <= ls2; i++ {
		if ls1%i == 0 && ls2%i == 0 {
			common = append(common, i)
		}
	}
	var check func(int) bool
	check = func(windowSize int) bool {
		// 1,1,1,1
		target := str1[:windowSize]
		start := windowSize
		for ; start <= ls1-windowSize; start += windowSize {
			end := start + windowSize
			if str1[start:end] != target {
				return false
			}
		}
		target1 := str2[:windowSize]
		if target1 != target {
			return false
		}
		start = windowSize
		for ; start <= ls2-windowSize; start += windowSize {
			end := start + windowSize
			if str2[start:end] != target1 {
				return false
			}
		}
		return true
	}
	for i := len(common) - 1; i >= 0; i-- {
		windowSize := common[i]
		if check(windowSize) {
			return str1[:windowSize]
		}
	}
	return ""
}
