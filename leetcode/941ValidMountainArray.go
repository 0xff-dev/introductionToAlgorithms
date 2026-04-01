package leetcode

func validMountainArray(arr []int) bool {
	size := len(arr)
	if size < 3 {
		return false
	}
	left, right := 0, size-1
	pre := -1
	for ; left < size && arr[left] > pre; left++ {
		pre = arr[left]
	}
	if left == size {
		return false
	}
	left--
	pre = -1
	for ; right >= 0 && arr[right] > pre; right-- {
		pre = arr[right]
	}
	if right == -1 {
		return false
	}
	right++
	return left == right
}
