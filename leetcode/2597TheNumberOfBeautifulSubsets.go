package leetcode

import (
	"math"
)

// 2,3,4,5,9,10
func beautifulSubsets(nums []int, k int) int {
	return countBeautifulSubsets(nums, k, 0, 0)
}

func countBeautifulSubsets(nums []int, difference, index, mask int) int {
	if index == len(nums) {
		if mask > 0 {
			return 1
		}
		return 0
	}

	isBeautiful := true

	for j := 0; j < index && isBeautiful; j++ {
		isBeautiful = ((1<<j)&mask) == 0 || math.Abs(float64(nums[j]-nums[index])) != float64(difference)
	}

	skip := countBeautifulSubsets(nums, difference, index+1, mask)
	var take int
	if isBeautiful {
		take = countBeautifulSubsets(nums, difference, index+1, mask+(1<<index))
	} else {
		take = 0
	}

	return skip + take
}
