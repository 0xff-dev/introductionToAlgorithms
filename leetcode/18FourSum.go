package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	sum := make([]map[int]struct{}, l)
	sum[l-1] = map[int]struct{}{
		nums[l-1]: {},
	}
	for i := l - 2; i >= 0; i-- {
		now := map[int]struct{}{
			nums[i]: {},
		}
		for k := range sum[i+1] {
			now[k] = struct{}{}
		}
		sum[i] = now
	}
	ans := make([][]int, 0)
	for i := 0; i < l-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		nf := true
		for next := i + 1; next < l-2; next++ {
			if !nf && nums[next] == nums[next-1] {
				continue
			}
			nf = false
			kf := true
			for k := next + 1; k < l-1; k++ {
				if !kf && nums[k] == nums[k-1] {
					continue
				}
				kf = false
				left := target - nums[k] - nums[next] - nums[i]
				if _, ok := sum[k+1][left]; ok {
					ans = append(ans, []int{nums[i], nums[next], nums[k], left})
				}
			}
		}
	}
	return ans
}
