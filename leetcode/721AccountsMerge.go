package leetcode

import (
	"sort"
)

/*
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
*/

type unionFind721 struct {
	father []int
}

func (u *unionFind721) find(x int) int {
	if u.father[x] != x {
		u.father[x] = u.find(u.father[x])
	}
	return u.father[x]
}
func (u *unionFind721) union(x, y int) {
	fx := u.find(x)
	fy := u.find(y)
	if fx < fy {
		u.father[fy] = fx
	} else {
		u.father[fx] = fy
	}

}
func accountsMerge(accounts [][]string) [][]string {
	u := unionFind721{father: make([]int, len(accounts))}
	for i := 0; i < len(accounts); i++ {
		u.father[i] = i
	}
	email2index := make(map[string]int)
	tmp := make([]map[string]struct{}, len(accounts))
	for idx, emails := range accounts {
		temp := map[string]struct{}{}
		for i := 1; i < len(emails); i++ {
			temp[emails[i]] = struct{}{}
			if targetIndex, ok := email2index[emails[i]]; ok {
				u.union(idx, targetIndex)
				continue
			}
			email2index[emails[i]] = idx
		}
		tmp[idx] = temp
	}
	for i := range accounts {
		f := u.find(i)
		if f != i {
			for key := range tmp[i] {
				tmp[f][key] = struct{}{}
			}
		}
	}
	result := make([][]string, 0)
	for i := range accounts {
		f := u.find(i)
		if f == i {
			result = append(result, []string{accounts[i][0]})
			next := make([]string, 0)
			for key := range tmp[i] {
				next = append(next, key)
			}
			sort.Strings(next)
			idx := len(result) - 1
			result[idx] = append(result[idx], next...)
		}
	}
	return result
}
