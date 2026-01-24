package leetcode

func centeredSubarrays(nums []int) int {
	ret := len(nums)
	l := ret
	sum := make([]int, l)
	sum[0] = nums[0]
	for i := 1; i < l; i++ {
		sum[i] = nums[i] + sum[i-1]
	}
	var cur int
	for i := 0; i < l-1; i++ {
		exist := map[int]struct{}{
			nums[i]: struct{}{},
		}
		for j := i + 1; j < l; j++ {
			exist[nums[j]] = struct{}{}
			cur = sum[j] - sum[i] + nums[i]
			if _, ok := exist[cur]; ok {
				ret++
			}
		}
	}
	return ret
}
