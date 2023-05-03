package leetcode

func findDifference(nums1 []int, nums2 []int) [][]int {
	ans := make([][]int, 2)
	nn1 := make([]bool, 2001)
	nn2 := make([]bool, 2001)
	for _, n := range nums1 {
		nn1[1000-n] = true
	}
	for _, n := range nums2 {
		nn2[1000-n] = true
	}
	for i := 0; i < 2001; i++ {
		if !nn1[i] && nn2[i] {
			ans[1] = append(ans[1], 1000-i)
		}
		if nn1[i] && !nn2[i] {
			ans[0] = append(ans[0], 1000-i)
		}
	}
	return ans
}
