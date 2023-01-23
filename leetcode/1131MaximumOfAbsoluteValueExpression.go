package leetcode

// 没想到更好的算法
func maxAbsValExpr(arr1 []int, arr2 []int) int {
	ans := 0
	l := len(arr1)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			diff1 := arr1[j] - arr1[i]
			diff2 := arr2[j] - arr2[i]
			if diff1 < 0 {
				diff1 = -diff1
			}
			if diff2 < 0 {
				diff2 = -diff2
			}
			d := j - i
			if d < 0 {
				d = -d
			}
			if r := diff1 + diff2 + d; r > ans {
				ans = r
			}
		}
	}
	return ans
}
