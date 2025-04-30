package leetcode

func findNumbers(nums []int) int {
	var isEven func(int) bool
	isEven = func(n int) bool {
		c := 0
		for n > 0 {
			n /= 10
			c++
		}
		return c&1 == 0
	}
	ans := 0
	for _, n := range nums {
		if isEven(n) {
			ans++
		}
	}
	return ans
}
