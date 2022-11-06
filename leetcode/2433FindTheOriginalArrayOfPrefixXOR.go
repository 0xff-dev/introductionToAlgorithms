package leetcode

func findArray(pref []int) []int {

	for idx := len(pref) - 1; idx > 0; idx-- {
		pref[idx] = pref[idx] ^ pref[idx-1]
	}

	return pref
}
