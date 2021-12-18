package leetcode

func removeDuplicatesII(nums []int) int {

	cmpObj, repeatCount := nums[0], 1
	fillIdx := 0
	for idx := 1; idx < len(nums); idx++ {
		if nums[idx] == cmpObj {
			repeatCount++
			continue
		}

		cmpObj = nums[idx]

		if repeatCount > 2 {
			repeatCount = 2
		}
		preObj := nums[idx-1]
		for ; repeatCount > 0; repeatCount, fillIdx = repeatCount-1, fillIdx+1 {
			nums[fillIdx] = preObj
		}
		repeatCount = 1
	}
	if repeatCount > 2 {
		repeatCount = 2
	}

	preObj := nums[len(nums)-1]
	for ; repeatCount > 0; repeatCount, fillIdx = repeatCount-1, fillIdx+1 {
		nums[fillIdx] = preObj
	}

	return fillIdx
}

func removeDuplicates1(nums []int) int {
	idx, repeat := 0, 1
	for walker := 1; walker < len(nums); walker++ {
		if nums[walker] == nums[idx] {
			if repeat != 2 {
				idx, repeat = idx+1, repeat+1
				nums[idx] = nums[walker]
			}
			continue
		}
		idx, repeat = idx+1, 1
		nums[idx] = nums[walker]
	}
	return idx+1
}