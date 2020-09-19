package sort

func InsertSort(nums []int) {
	for index := 1; index < len(nums); index++ {
		lowIndex, num := index-1, nums[index]
		for lowIndex >= 0 && nums[lowIndex] >= nums[index] {
			lowIndex--
		}
		for pre := index - 1; pre > lowIndex; pre-- {
			nums[pre+1] = nums[pre]
		}
		nums[lowIndex+1] = num
	}
}
