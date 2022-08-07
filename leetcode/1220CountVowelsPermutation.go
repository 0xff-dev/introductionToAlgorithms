package leetcode

const mod1220 = 1000000007

// 最常规思路 递归硬计算
//var followedBy = map[byte][]byte{
//	' ': {'a', 'e', 'i', 'o', 'u'},
//	'a': {'e'},
//	'e': {'a', 'i'},
//	'i': {'a', 'e', 'o', 'u'},
//	'o': {'i', 'u'},
//	'u': {'a'},
//}
//
//func countVowelPermutation(n int) int {
//	ans := 0
//	countVowelPermutationDFS(0, n, &ans, ' ')
//	return ans
//}
//
//func countVowelPermutationDFS(idx, n int, ans *int, preByte byte) {
//	if idx == n {
//		*ans = (*ans + 1) % mod1220
//		return
//	}
//	for _, b := range followedBy[preByte] {
//		countVowelPermutationDFS(idx+1, n, ans, b)
//	}
//}

// 类似黑白房那个动作，
// 结尾两个是黑黑只能接白的那种感觉 dp 确定是如何结尾的
func countVowelPermutation(n int) int {
	a, e, i, o, u := 1, 1, 1, 1, 1

	//	'a': {'e'},
	//	'e': {'a', 'i'},
	//	'i': {'a', 'e', 'o', 'u'},
	//	'o': {'i', 'u'},
	//	'u': {'a'},
	for idx := 2; idx <= n; idx++ {
		// e, i, u后面可以接a
		a1 := (e + i + u) % mod1220 // x
		e1 := (a + i) % mod1220     //
		i1 := (e + o) % mod1220     //
		o1 := i % mod1220
		u1 := (i + o) % mod1220
		a, e, i, o, u = a1, e1, i1, o1, u1
	}
	return (a + e + i + o + u) % mod1220
}
