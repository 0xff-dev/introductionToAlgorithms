package leetcode

// 感觉dp可解决, 后续思考后补充
func findTargetSumWays(nums []int, target int) int {
	ans := 0
	var dfs func(int, int)
	dfs = func(index, sum int) {
		if index == len(nums) {
			if sum == target {
				ans++
			}
			return
		}
		dfs(index+1, sum-nums[index])
		dfs(index+1, sum+nums[index])
	}
	dfs(0, 0)
	return ans
}
