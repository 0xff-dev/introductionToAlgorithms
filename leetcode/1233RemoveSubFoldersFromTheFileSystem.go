package leetcode

import (
	"sort"
	"strings"
)

/*
type filePathNode struct {
	end  bool
	data map[string]*filePathNode
}

func insertFolder(root *filePathNode, path string) {
	walker := root
	paths := strings.Split(path, "/")
	for i := 1; i < len(paths); i++ {
		p := paths[i]
		_, ok := walker.data[p]
		if !ok {
			walker.data[p] = &filePathNode{end: false, data: make(map[string]*filePathNode)}
		}
		if i == len(paths)-1 {
			walker.data[p].end = true
			continue
		}
		if walker.data[p].end {
			continue
		}
		walker = walker.data[p]
	}
}
func searchFolder(root *filePathNode, path string) []string {
	ans := []string{}
	if root.end {
		ans = append(ans, path)
		return ans
	}
	for part, next := range root.data {
		ans = append(ans, searchFolder(next, path+"/"+part)...)
	}
	return ans
}
func removeSubfolders(folder []string) []string {
	root := &filePathNode{end: false, data: make(map[string]*filePathNode)}
	for _, str := range folder {
		insertFolder(root, str)
	}
	return searchFolder(root, "")
}
*/

type filePathNode struct {
	end  bool
	data map[string]*filePathNode
}

func insertFolder(root *filePathNode, path string) bool {
	walker := root
	paths := strings.Split(path, "/")
	for i := 1; i < len(paths); i++ {
		p := paths[i]
		_, ok := walker.data[p]
		if !ok {
			walker.data[p] = &filePathNode{end: false, data: make(map[string]*filePathNode)}
		}
		if i == len(paths)-1 {
			walker.data[p].end = true
			continue
		}
		if walker.data[p].end {
			return false
		}
		walker = walker.data[p]
	}
	return true
}

func removeSubfolders(folder []string) []string {
	sort.Slice(folder, func(i, j int) bool {
		return len(folder[i]) < len(folder[j])
	})

	root := &filePathNode{end: false, data: make(map[string]*filePathNode)}
	ans := make([]string, 0)
	for _, str := range folder {
		if insertFolder(root, str) {
			ans = append(ans, str)
		}
	}
	return ans
}
