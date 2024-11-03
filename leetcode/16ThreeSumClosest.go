package leetcode

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	l := len(nums)
	first := true
	ans := 0
	diff := 0
	var d1 int
	for i := 0; i < l-2; i++ {
		for j := i + 1; j < l-1; j++ {
			sum := nums[i] + nums[j]
			start := j + 1
			ll := l - start
			idx := sort.Search(ll, func(k int) bool {
				return nums[start+k]+sum >= target
			})
			if idx == ll {
				sum += nums[l-1]
				d1 = sum - target
				if d1 < 0 {
					d1 = -d1
				}
			} else if idx == 0 {
				sum += nums[start]
				d1 = sum - target
				if d1 < 0 {
					d1 = -d1
				}
			} else {
				a := sum + nums[start+idx] - target
				if a < 0 {
					a = -a
				}
				b := sum + nums[start+idx-1] - target
				if b < 0 {
					b = -b
				}
				if a < b {
					sum += nums[start+idx]
					d1 = a
				} else {
					sum += nums[start+idx-1]
					d1 = b
				}
			}
			if first {
				ans = sum
				diff = d1
				first = false
				continue
			}
			if diff > d1 {
				ans = sum
				diff = d1
			}
		}
	}
	return ans
}
