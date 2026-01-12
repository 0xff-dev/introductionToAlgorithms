package leetcode

import (
	"strconv"
)

func getFolderNames(names []string) []string {
	l := len(names)
	ret := make([]string, l)
	cache := make(map[string]int)
	var tmpName, suffix string
	for idx, name := range names {
		if _, ok := cache[name]; !ok {
			cache[name] = 1
			ret[idx] = name
			continue
		}
		nextNumber := cache[name]
		for {
			suffix = strconv.Itoa(nextNumber)
			tmpName = name + "(" + suffix + ")"
			_, ok := cache[tmpName]
			if ok {
				nextNumber++
				continue
			}
			ret[idx] = tmpName
			cache[tmpName] = 1
			nextNumber++
			cache[name] = nextNumber
			break
		}
	}
	return ret
}
