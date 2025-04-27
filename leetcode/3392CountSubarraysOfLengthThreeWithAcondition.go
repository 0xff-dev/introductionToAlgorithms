package leetcode

func countSubarrays3392(nums []int) int {
	ans := 0
	for i := 0; i < len(nums)-2; i++ {
		a := nums[i]
		b := nums[i+1]
		c := nums[i+2]
		if (a+c)*2 == b {
			ans++
		}
	}
	return ans
}
