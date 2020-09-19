package sort

func merge(nums []int, start, end int) {
	mid := start + (end-start)/2
	l, r := start, mid+1
	tmpNums := make([]int, 0)
	for l <= mid && r <= end {
		if nums[l] < nums[r] {
			tmpNums = append(tmpNums, nums[l])
			l++
		} else {
			tmpNums = append(tmpNums, nums[r])
			r++
		}
	}
	for ; l <= mid; l++ {
		tmpNums = append(tmpNums, nums[l])
	}
	for ; r <= end; r++ {
		tmpNums = append(tmpNums, nums[r])
	}
	for k := 0; k < len(tmpNums); k++ {
		nums[k+start] = tmpNums[k]
	}
}

func MergeSort(nums []int, start, end int) {
	if start < end {
		mid := start + (end-start)/2
		MergeSort(nums, start, mid)
		MergeSort(nums, mid+1, end)
		merge(nums, start, end)
	}
}
