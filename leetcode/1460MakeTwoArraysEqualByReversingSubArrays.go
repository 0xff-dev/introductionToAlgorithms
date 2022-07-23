package leetcode

/*
1 2 3 4 5
5 4 3 2 1
*/
func canBeEqual(target []int, arr []int) bool {
	check := make([]int, 1001)
	for idx := 0; idx < len(target); idx++ {
		check[target[idx]]++
	}
	for idx := 0; idx < len(target); idx++ {
		check[arr[idx]]--
	}
	for idx := 0; idx < 1001; idx++ {
		if check[idx] != 0 {
			return false
		}
	}
	return true
}
