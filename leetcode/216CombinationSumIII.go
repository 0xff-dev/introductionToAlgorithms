package leetcode

func combinationSum3(k, n int) [][]int {
	result := make([][]int, 0)
	combinationSum3Aux(0, 1, k, n, &result, []int{})
	return result
}

func combinationSum3Aux(sum, num, k, target int, result *[][]int, paths []int) {
	if num > 9 {
		if len(paths) == k && sum == target {
			dst := make([]int, len(paths))
			copy(dst, paths)
			*result = append(*result, dst)
		}
		return
	}
	combinationSum3Aux(sum, num+1, k, target, result, paths)
	combinationSum3Aux(sum+num, num+1, k, target, result, append(paths, num))
}
