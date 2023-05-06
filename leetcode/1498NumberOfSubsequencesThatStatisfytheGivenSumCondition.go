package leetcode

import (
	"sort"
)

const mod1498 = 1000000007

/*
type mm1498 struct {
	m, n int
}

func numSubseq(nums []int, target int) int {
	ans := 0
	length := len(nums)
	mm := make([][]mm1498, length)
	for i := 0; i < length; i++ {
		mm[i] = make([]mm1498, length)
	}
	// 长度为1的情况
	for i := 0; i < length; i++ {
		mm[i][i] = mm1498{i, i}
	}

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[mm[i][j-1].n] < nums[j] {
				mm[i][j].n = mm[i][j-1].n
			} else {
				mm[i][j].n = j
			}

			if nums[mm[i][j-1].m] < nums[j] {
				mm[i][j].m = j
			} else {
				mm[i][j].m = mm[i][j-1].m
			}
		}
	}
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if nums[mm[i][j].m]+nums[mm[i][j].n] <= target {
				ans = (ans + 1) % mod1498
			}
		}
	}
	return ans
}
*/

/*
func numSubseq(nums []int, target int) int {
	ans := 0
	for l := 1; l <= len(nums); l++ {
		subset(nums, 0, l, target, -1, -1, &ans)
	}
	return ans
}

func subset(nums []int, index, length, target, m, n int, ans *int) {
	if length == 0 {
		if m+n <= target {
			*ans = (*ans + 1) % mod1498
		}
		return
	}
	if index >= len(nums) {
		return
	}
	subset(nums, index+1, length, target, m, n, ans)
	if m == -1 || nums[index] > m {
		m = nums[index]
	}
	if n == -1 || nums[index] < n {
		n = nums[index]
	}
	subset(nums, index+1, length-1, target, m, n, ans)
}
*/

func numSubseq(nums []int, target int) int {
	ans := 0
	sort.Ints(nums)
	power := make([]int, len(nums))
	power[0] = 1
	for i := 1; i < len(nums); i++ {
		power[i] = (power[i-1] * 2) % mod1498
	}
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left]+nums[right] <= target {
			ans += power[right-left]
			ans %= mod1498
			left++
			continue
		}
		right--
	}
	return ans
}
