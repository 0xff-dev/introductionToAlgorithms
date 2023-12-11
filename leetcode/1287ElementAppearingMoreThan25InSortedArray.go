package leetcode

func findSpecialInteger(arr []int) int {
	count := 1
	l := len(arr)
	need := l / 4
	addOne := false
	if l%4 != 0 {
		need++
		addOne = true
	}
	ans := -1
	i := 1
	for ; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			count++
			continue
		}
		if (!addOne && count > need) || (addOne && count >= need) {
			ans = arr[i-1]
			break
		}
		count = 1
	}
	if (!addOne && count > need) || (addOne && count >= need) {
		ans = arr[i-1]
	}
	return ans
}
