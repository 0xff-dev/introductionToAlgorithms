package leetcode

/*
func canPartition(nums []int) bool {
    sum := 0
    // 感觉又是计算奇数与偶数的个数的问题呢
    for _, n := range nums {
        sum += n
    }
    if sum & 1 == 1 {
        return false
    }
    return canPartitionDfs(nums, 0, 0, sum/2)
}

func canPartitionDfs(nums []int, idx, sum, target int) bool {
    if idx == len(nums) {
        return sum == target
    }

    if(canPartitionDfs(nums, idx+1, sum, target)) {
        return true
    }

    if (canPartitionDfs(nums, idx+1, sum+nums[idx], target)) {
        return true
    }
    return false
}
*/

func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum&1 == 1 {
		return false
	}

	sum /= 2
	dp := make([]bool, sum+1)
	dp[0] = true
	for _, n := range nums {
		for walker := sum; walker >= n; walker-- {
			dp[walker] = dp[walker] || dp[walker-n]
		}
	}
	return dp[sum]
}
