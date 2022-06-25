package leetcode

func checkPossibility(nums []int) bool {
	diff := 0
	cmpObj := nums[0]
	for idx := 1; idx < len(nums); idx++ {
		if nums[idx] >= cmpObj {
			cmpObj = nums[idx]
			continue
		}

		// 1, 2, 3, 5, 2, 6
		// 当比较2，5的时候，发现2比3小，前面都已经排序了，所以只能调整2变成5, 这个时候让cmpObj保持不变即可
		// 如果比较的是 1, 2, 3, 5, 4, 6 这个时候发现4比3所以可以直接将调整为4，变更cmpObj=4
		// 但是记得最大变更次数是1
		if idx == 1 || idx >= 2 && nums[idx] >= nums[idx-2] {
			cmpObj = nums[idx]
		}
		diff++
		if diff > 1 {
			return false
		}
	}
	return true
}
