package leetcode

// 1 - A
// 2 - A
// 3 -
func PredictTheWinner(nums []int) bool {
	n := len(nums)
	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, n)
		for j := 0; j < n; j++ {
			cache[i][j] = -1
		}
	}
	var search func(int, int) int
	search = func(left, right int) int {
		if cache[left][right] != -1 {
			return cache[left][right]
		}
		if left == right {
			return nums[left]
		}
		l := nums[left] - search(left+1, right)
		r := nums[right] - search(left, right-1)
		cache[left][right] = l
		if r > l {
			cache[left][right] = r
		}
		return cache[left][right]
	}

	return search(0, n-1) >= 0
}
