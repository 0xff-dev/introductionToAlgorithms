package sort

func BubbleSort(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}
	for cmp := 0; cmp < length-1; cmp++ {
		for i := 0; i < length-1-cmp; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
}
