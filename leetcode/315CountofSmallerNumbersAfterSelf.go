package leetcode

func countSmaller(nums []int) []int {
	lessCount, sortNums := make([]int, len(nums)), make([]int, len(nums))
	for idx, num := range nums {
		sortNums[idx] = num
	}
	mergeSortAux(sortNums, lessCount, 0, len(nums)-1)
	result := make([]int, len(nums))
	/// 这里感觉可以优化
	tmp := make(map[int][]int)
	for idx, item := range sortNums {
		if _, ok := tmp[item]; !ok {
			tmp[item] = []int{}
		}
		tmp[item] = append(tmp[item], lessCount[idx])
	}
	for idx, item := range nums {
		r := tmp[item]
		result[idx] = r[0]
		if len(r) > 0 {
			tmp[item] = tmp[item][1:]
		}
	}
	return result
}

func mergeAux(nums, lessCount []int, start, end int) {
	mid := start + (end-start)/2
	tmp, tmpLess := make([]int, end-start+1), make([]int, end-start+1)
	left, right := start, mid+1
	base, k := 0, 0
	for left <= mid && right <= end {
		if nums[left] <= nums[right] {
			tmp[k] = nums[left]
			tmpLess[k] = lessCount[left] + base
			left++
		} else {
			tmp[k] = nums[right]
			tmpLess[k] = lessCount[right]
			lessCount[right] = 0
			right, base = right+1, base+1
		}
		k++
	}
	for ; left <= mid; left, k = left+1, k+1 {
		tmp[k] = nums[left]
		tmpLess[k] = lessCount[left] + base
	}
	for ; right <= end; right, k = right+1, k+1 {
		tmp[k] = nums[right]
		tmpLess[k] = lessCount[right]
	}

	for idx, num := range tmp {
		nums[start+idx] = num
		lessCount[start+idx] = tmpLess[idx]
	}
}

func mergeSortAux(nums, lessCount []int, start, end int) {
	if start < end {
		mid := start + (end-start)/2
		mergeSortAux(nums, lessCount, start, mid)
		mergeSortAux(nums, lessCount, mid+1, end)
		mergeAux(nums, lessCount, start, end)
	}
}
