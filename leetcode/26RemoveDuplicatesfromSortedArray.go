package leetcode

//func removeDuplicates(nums []int) int {
//	length := len(nums)
//	index := length-1
//	for index > 0 {
//		pre := index-1
//		repeat := 0
//		for ;pre >= 0 && nums[pre] == nums[index]; pre-- {
//			repeat++
//		}
//
//		for i := index; i < length; i++ {
//			nums[i-repeat] = nums[i]
//		}
//		length = length-repeat
//		index = pre
//	}
//	return length
//}

func removeDuplicates(nums []int) int {
	index := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[index] {
			index++
			nums[index] = nums[i]
		}
	}
	return index + 1
}
