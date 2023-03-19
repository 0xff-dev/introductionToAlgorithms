package challengeprogrammingdatastructure

// s 是生序
func BinarySearch(s, t []int) []int {

	ans := make([]int, 0)
	for _, key := range t {
		if binarySearch(s, key) {
			ans = append(ans, key)
		}
	}
	return ans
}

func binarySearch(s []int, key int) bool {
	left, right := 0, len(s)-1
	for left <= right {
		mid := (right-left)/2 + left
		if s[mid] == key {
			return true
		}
		if s[mid] < key {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
