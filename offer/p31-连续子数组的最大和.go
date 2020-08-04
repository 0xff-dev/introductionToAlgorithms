package offer

import "fmt"

func MaxSumInSubArray(nums []int) int {
	maxSum, start, end, sum := 0, 0, 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > maxSum {
			maxSum = sum
			end = i
		} else if sum < 0 {
			sum = 0
			start, end = i+1, i+1
		}
	}
	fmt.Println("res: ", nums[start:end+1])
	return maxSum
}
