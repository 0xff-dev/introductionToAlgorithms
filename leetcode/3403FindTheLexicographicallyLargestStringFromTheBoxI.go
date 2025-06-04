package leetcode

func answerString(word string, numFriends int) string {
	length := len(word)
	ans := ""
	if length < numFriends {
		return ans
	}
	if numFriends == 1 {
		return word
	}

	indies := [26][]int{}
	for i, b := range word {
		indies[b-'a'] = append(indies[b-'a'], i)
	}
	i := 25
	for ; i >= 0 && len(indies[i]) == 0; i-- {
	}

	for _, start := range indies[i] {
		tmp := ""
		left := start
		if left >= numFriends {
			tmp = word[start:]
		} else {
			if end := length - (numFriends - 1 - left); end >= start {
				tmp = word[start:end]
			}
		}
		if tmp > ans {
			ans = tmp
		}
	}
	return ans
}
