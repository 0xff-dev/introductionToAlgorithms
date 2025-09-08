package leetcode

func ok1317(n int) bool {
	for n > 0 {
		mod := n % 10
		if mod == 0 {
			return false
		}
		n /= 10
	}
	return true
}

func getNoZeroIntegers(n int) []int {
	nums := make(map[int]struct{})
	for i := 1; i <= 10000; i++ {
		if ok1317(i) {
			nums[i] = struct{}{}
		}
	}
	var diff int
	for k := range nums {
		diff = n - k
		if _, ok := nums[diff]; ok {
			return []int{k, diff}
		}
	}
	return []int{}
}
