package leetcode

import "fmt"

func minDeletions1647(s string) int {
	letters := make([]int, 26)
	for _, b := range []byte(s) {
		letters[b-'a']++
	}
	occurrences := make(map[int]struct{})
	// 1 1, 2, 2, 2, 3, 3, 4, 4, 5  需要4次
	// 2    3        2     2     1
	// 1   0         0     2     1  4 次
	// 1, 1, 2, 2, 2, 3, 3, 2
	// 2     3  2
	/// 0     0  1
	// 2:3, 3:1 1:1
	ans := 0
	//for idx := 0; idx < 26; idx++ {
	//	if letters[idx] == 0 {
	//		continue
	//	}
	//	fmt.Printf("%c = %d\n", 'a'+idx, letters[idx])
	//}
	for idx := 0; idx < 26; idx++ {
		if letters[idx] == 0 {
			continue
		}
		r := letters[idx]
		if _, ok := occurrences[r]; !ok {
			occurrences[r] = struct{}{}
			continue
		}
		for r = r - 1; r >= 0; r-- {
			ans++
			if _, ok := occurrences[r]; !ok {
				occurrences[r] = struct{}{}
				break
			}
		}
	}

	fmt.Println()
	return ans
}
