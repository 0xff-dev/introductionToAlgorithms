package leetcode

/*
func wonderfulSubstrings(word string) int64 {
	// 10bit
	ans := int64(0)
	// 0 0 0 0 0 0 0 0 0 0
	// j i h g f e d c b a
	l := len(word)
	// 0 或者(n& n-1) == 0 就是满足要求的。
	state := make(map[int]int64)
	state[1<<int(word[0]-'a')] = 1
	ans++
	for i := 1; i < l; i++ {
		pos := 1 << int(word[i]-'a')
		next := make(map[int]int64)
		next[pos] = 1
		ans++
		//ab:1 b:1
		for s, c := range state {
			cur := pos ^ s
			next[cur] += c
			if cur == 0 && (cur&(cur-1) == 0) {
				ans += c
			}
		}
		state = next
		fmt.Printf("--- next state: %+v\n", state)
	}
	return ans
}
*/

func wonderfulSubstrings(word string) int64 {
	// 10bit
	ans := int64(0)
	l := len(word)
	// 0 或者(n^n-1) == 0 就是满足要求的。
	state := make(map[int]int64)
	state[0] = 1
	mask := 0
	for i := 0; i < l; i++ {
		mask ^= (1 << int(word[i]-'a'))
		c := state[mask]
		ans += c
		state[mask] = c + 1
		for i := 0; i < 10; i++ {
			ans += state[mask^(1<<i)]
		}
	}
	return ans
}
