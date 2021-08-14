package leetcode

func countGoodTriplets(arr []int, a, b, c int) int {
	count := 0

	length := len(arr)
	if length < 3 {
		return 0
	}

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if !(abs1(arr[i]-arr[j]) <= a) {
				continue
			}
			for k := j + 1; k < length; k++ {
				a1, b1 := abs1(arr[j]-arr[k]), abs1(arr[i]-arr[k])
				if a1 <= b && b1 <= c {
					count++
				}
			}
		}
	}

	return count
}

func abs1(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
