package leetcode

import "sort"

type storeItem struct {
	timestamp int
	val       string
}
type TimeMap struct {
	data map[string][]storeItem
}

func Constructor981() TimeMap {
	return TimeMap{
		data: make(map[string][]storeItem),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.data[key]; !ok {
		this.data[key] = make([]storeItem, 0)
	}
	this.data[key] = append(this.data[key], storeItem{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	v, ok := this.data[key]
	if !ok {
		return ""
	}
	idx := sort.Search(len(v), func(i int) bool {
		return v[i].timestamp > timestamp
	})
	if idx == 0 {
		return ""
	}
	return v[idx-1].val
}
