package leetcode

func kidsWithCandies(candies []int, extraCandies int) []bool {
	length := len(candies)
	l, r := make([]int, length), make([]int, length)
	l[0], r[length-1] = candies[0], candies[length-1]
	for i := 1; i < length; i++ {
		l[i] = l[i-1]
		if candies[i] > l[i] {
			l[i] = candies[i]
		}
	}
	for i := length - 2; i >= 0; i-- {
		r[i] = r[i+1]
		if candies[i] > r[i] {
			r[i] = candies[i]
		}
	}
	ans := make([]bool, length)
	for i := 0; i < length; i++ {
		ll := true
		if i >= 1 && candies[i]+extraCandies < l[i-1] {
			ll = false
		}
		rr := true
		if i+1 < length && candies[i]+extraCandies < r[i+1] {
			rr = false
		}
		ans[i] = ll && rr
	}
	return ans
}
