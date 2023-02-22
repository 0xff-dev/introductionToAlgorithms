package leetcode

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	ans := 0
	for i := 0; i < len(arr1); i++ {
		j := 0
		for ; j < len(arr2); j++ {
			diff := arr1[i] - arr2[j]
			if diff < 0 {
				diff = ^diff + 1
			}
			if diff <= d {
				break
			}
		}
		if j == len(arr2) {
			ans++
		}
	}
	return ans
}
