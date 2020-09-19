package sort

func HeapSort(nums []int)  {
	build(nums)
	length := len(nums)
	for length != 0 {
		nums[0], nums[length-1] = nums[length-1], nums[0]
		length--
		maintain(nums, length, 0)
	}
}

func build(nums []int) {
	for index := len(nums)/2; index >= 0; index-- {
		maintain(nums, len(nums), index)
	}
}

func maintain(nums []int, heapSize, index int){
	left, right := index*2, index*2+1
	maintainIndex := index
	if left < heapSize && nums[left] > nums[maintainIndex] {
		maintainIndex = left
	}
	if right < heapSize && nums[right] > nums[maintainIndex] {
		maintainIndex = right
	}
	if maintainIndex != index {
		nums[maintainIndex], nums[index] = nums[index], nums[maintainIndex]
		maintain(nums, heapSize, maintainIndex)
	}
}

