package leetcode

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	sum := 0
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			left := num / i
			sum += i
			if left != i {
				sum += left
			}
		}
	}
	return sum+1 == num
}
