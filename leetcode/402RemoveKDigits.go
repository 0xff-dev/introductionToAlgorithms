package leetcode

/*
func removeKdigits(num string, k int) string {
	// 1234 3
	//
	// dfs + cache
	n := len(num)
	if k == n {
		return "0"
	}
	cache := make(map[string]map[int]string)
	var dfs func(string, int) string
	dfs = func(cur string, left int) string {
		i := 0
		for ; i < len(cur); i++ {
			if cur[i] != '0' {
				break
			}
		}
		cur = cur[i:]

		if left == 0 {
			return cur
		}
		if v, ok := cache[cur]; ok {
			if v1, ok1 := v[left]; ok1 {
				return v1
			}
		}

		ans := cur
		for j := 0; j < len(cur); j++ {
			// 10200
			next := cur[:j] + cur[j+1:]
			tmp := dfs(next, left-1)
			//fmt.Printf("------ next %s --- k %d tmp: %s\n", next, left, tmp)
			if len(tmp) < len(ans) || (len(tmp) == len(ans) && ans > tmp) {
				ans = tmp
			}
			//fmt.Printf("--- got ans %s\n", ans)
		}
		if _, ok := cache[cur]; !ok {
			cache[cur] = make(map[int]string)
		}
		cache[cur][left] = ans
		return ans
	}

	return dfs(num, k)
}
*/

func removeKdigits(num string, k int) string {
	n := len(num)
	if n == k {
		return "0"
	}
	stack := make([]byte, 0)
	for idx := 0; idx < n; idx++ {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[idx] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[idx])
	}
	stack = stack[:len(stack)-k]
	i := 0
	for ; i < len(stack) && stack[i] == '0'; i++ {
	}
	if i == len(stack) {
		return "0"
	}
	return string(stack[i:])
}
