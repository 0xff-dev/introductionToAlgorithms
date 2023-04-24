package challengeprogrammingdatastructure

/*
首先，除2以外所有的偶数都不是质数，这样就能将复杂度减少一半，
然后在检查x时， 由于x不可能被大于x/2 的整数整除，这就又减少了一半复杂度。
但这些小技巧并不能撼动该 算法复杂度为O(x)的本质
合数 x 拥有满足 p*p <= x 的质因数 p
*/
// 埃拉托色尼筛选法, 将质数的倍数都挑出去。
func isPrime(x int) bool {
	if x == 2 {
		return true
	}

	if x < 2 || x&1 == 0 {
		return false
	}
	i := 3
	for i*i <= x {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func PrimeNumbers(nums []int) int {
	ans := 0
	for _, n := range nums {
		if isPrime(n) {
			ans++
		}
	}
	return ans
}
