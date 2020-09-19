package sort

func QuickSort(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	aim, pre, walker := nums[0], 0, 1
	for walker < len(nums) {
		if nums[walker] < aim {
			pre++
			nums[walker], nums[pre] = nums[pre], nums[walker]
		}
		walker++
	}
	nums[0], nums[pre] = nums[pre], nums[0]
	QuickSort(nums[:pre])
	QuickSort(nums[pre+1:])
}
