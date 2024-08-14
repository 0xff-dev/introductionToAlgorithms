package leetcode

import (
	"sort"
)

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	// 通常情况下 那基本就是n^2 + nlog(n)了，计算出全部，然后重新排列
	// 对于条件i < j, 就是选择(i, j)或者(j, i)的问题, 可以直接忽略了
	// 如果完全计算，肯定是不行的，10^8 如果选择最后几个元素，内存都要爆炸
	l := len(nums)
	var count func(int) int
	/*
	  	count = func(target int) (int, int) {
	  		s, e := 0, 1
	  		// 1, 1, 6
	  		less, equal := 0, 0
	  		for ; e < l; e++ {
	  			diff := nums[e] - nums[s]
	  			if diff < target {
	  				fmt.Printf("add [%d, %d]\n", e, s)
	  				less++
	  				continue
	  			}
	  			if diff == target {
	  				equal++
	  				continue
	  			}
	  			s++
	  			for ; s < e; s++ {
	  				diff = nums[e] - nums[s]
	  				if diff <= target {
	  					break
	  				}
	  			}

	  			if s != e {
	  				if diff < target {
	  					less++
	  				} else {
	  					equal++
	  				}
	  			}
	  		}
	  		s++
	  		for ; s < e-1; s++ {
	  			diff := nums[e-1] - nums[s]

	  			if diff < target {
	  				less++
	  				continue
	  			}
	  			equal++
	  		}
	  		return less, equal
	  	}
	  麻了
	*/
	count = func(target int) int {
		ans := 0
		// 4, 62, 100
		for i := 1; i < l; i++ {
			idx := sort.Search(i, func(ii int) bool {
				return nums[i]-nums[ii] <= target
			})
			ans += i - idx
		}
		return ans
	}
	a, b := 0, nums[l-1]-nums[0]
	// 查看小于等于当前值的有多少个
	// 4, 62, 100
	// 38,58,96,
	// 0, 5, 5
	// 0 ~ 5
	// 3 ~ 5
	// 4 ~ 5
	for a < b {
		mid := (a + b) / 2
		c := count(mid)
		if c < k {
			a = mid + 1
			continue
		}
		b = mid
	}
	return a
}
