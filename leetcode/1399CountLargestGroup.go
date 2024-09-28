package leetcode

func countLargestGroup(n int) int {
	cm := make(map[int]int)
	count := 0
	ans := 0
	for i := 1; i <= n; i++ {
		tmp := 0
		j := i
		for j > 0 {
			tmp += j % 10
			j /= 10
		}
		cm[tmp]++
		if cm[tmp] > count {
			count = cm[tmp]
			ans = 0
		}
		if cm[tmp] == count {
			ans++
		}
	}
	return ans
}
