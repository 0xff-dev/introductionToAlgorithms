package leetcode

func quickSort912(nums []int, left, right int) {
	if right <= left {
		return
	}
	target := nums[left]
	pre := left
	for i := left + 1; i <= right; i++ {
		if nums[i] < target {
			pre++
			nums[i], nums[pre] = nums[pre], nums[i]
		}
	}
	nums[left], nums[pre] = nums[pre], nums[left]
	quickSort912(nums, left, pre-1)
	quickSort912(nums, pre+1, right)
}

func merge912(nums []int, left, right int) {
	mid := (right-left)/2 + left
	i, j := left, mid+1
	k := 0
	store := make([]int, right-left+1)
	for ; i <= mid && j <= right; k++ {
		if nums[i] < nums[j] {
			store[k] = nums[i]
			i++
		} else {
			store[k] = nums[j]
			j++
		}
	}
	for ; i <= mid; i, k = i+1, k+1 {
		store[k] = nums[i]
	}
	for ; j <= right; j, k = j+1, k+1 {
		store[k] = nums[j]
	}
	for x := 0; x < k; x++ {
		nums[left+x] = store[x]
	}
}

func mergeSort912(nums []int, left, right int) {
	if left < right {
		mid := (right-left)/2 + left
		mergeSort912(nums, left, mid)
		mergeSort912(nums, mid+1, right)
		merge912(nums, left, right)
	}
}

func sortArray(nums []int) []int {
	//quickSort912(nums, 0, len(nums)-1)
	mergeSort912(nums, 0, len(nums)-1)
	return nums
}
