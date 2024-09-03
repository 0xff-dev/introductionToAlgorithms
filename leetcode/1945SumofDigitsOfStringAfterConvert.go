package leetcode

func getLucky(s string, k int) int {
	// s[i]-'a' + uint8(1)
	// 我估计是数学题
	sum := 0
	k--
	for _, b := range s {
		tmp := int(b-'a') + 1
		for tmp > 0 {
			sum += tmp % 10
			tmp /= 10
		}
	}
	for ; k > 0; k-- {
		tmp := 0
		// TOOD sum
		for sum > 0 {
			tmp += sum % 10
			sum /= 10
		}
		sum = tmp
	}

	return sum
}
