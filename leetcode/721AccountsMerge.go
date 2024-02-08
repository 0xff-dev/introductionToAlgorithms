package leetcode

import "sort"

type account struct {
	name   string
	emails map[string]struct{}
}

func accountsMerge(accounts [][]string) [][]string {
	emailBelong := make(map[string]int)
	ans := make([]account, 0)
	deleteIndies := make(map[int]struct{})
	for _, emails := range accounts {
		username := emails[0]
		found := []int{}
		minIndex := -1
		tmp := make(map[string]struct{})
		for i := 1; i < len(emails); i++ {
			tmp[emails[i]] = struct{}{}
			if targetIndex, ok := emailBelong[emails[i]]; ok {
				found = append(found, targetIndex)
				if minIndex == -1 || minIndex > targetIndex {
					minIndex = targetIndex
				}
			}
		}
		setIndex := 0
		if len(found) == 0 {
			ans = append(ans, account{name: username, emails: tmp})
			setIndex = len(ans) - 1
		} else {
			for _, idx := range found {
				if idx == minIndex {
					continue
				}
				for key := range ans[idx].emails {
					ans[minIndex].emails[key] = struct{}{}
				}
				deleteIndies[idx] = struct{}{}
			}
			for key := range tmp {
				ans[minIndex].emails[key] = struct{}{}
			}
			setIndex = minIndex
		}
		for key := range ans[setIndex].emails {
			emailBelong[key] = setIndex
		}
	}
	result := make([][]string, 0)
	for i := range ans {
		if _, ok := deleteIndies[i]; ok {
			continue
		}
		emails := make([]string, 0)
		for e := range ans[i].emails {
			emails = append(emails, e)
		}
		sort.Strings(emails)
		next := append([]string{ans[i].name}, emails...)
		result = append(result, next)
	}

	return result
}
