package leetcode

/*
func subarraysDivByK(nums []int, k int) int {
	length := len(nums)
	sum := make([]int, length)
	sum[0] = nums[0]
	pre := 0
	ans := 0
	for i := 0; i < length; i++ {
		if nums[i] % k == 0 {
			ans++
		}
		sum[i] = pre + nums[i]
		pre = sum[i]
	}

	// 1,2,3,4
	for l := 2; l <= length; l++ {
		pre := 0
		for start := l-1; start < length; start++ {
			diff := sum[start]-pre
			if diff % k == 0 {
				ans++
			}
			pre = sum[start+1-l]
		}
	}
	return ans
}
*/

func subarraysDivByK(nums []int, k int) int {
	modK := make([]int, k)
	modK[0] = 1
	preMod, res := 0, 0
	for i := 0; i < len(nums); i++ {
		preMod = (preMod + nums[i]%k + k) % k
		res += modK[preMod]
		modK[preMod]++
	}
	return res
}
