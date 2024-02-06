package leetcode

import (
	"strings"
)

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
