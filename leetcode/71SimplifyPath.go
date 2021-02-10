package leetcode

import (
	"bytes"
	"strings"
)

func simplifyPath(path string) string {
	if path == "" {
		return path
	}

	paths := strings.Split(path, "/")
	newPathFields := make([]string, len(paths))
	idx := 0
	for _, p := range paths {
		if p == "." || p == "" {
			continue
		}

		if p == ".." {
			if idx == 0 {
				continue
			}
			idx--
			continue
		}

		newPathFields[idx] = p
		idx++
	}
	if idx == 0 {
		return "/"
	}

	buf := bytes.NewBufferString("")
	if path[0] == '/' || paths[0] == ".." {
		buf.WriteByte('/')
	}
	for i := 0; i < idx-1; i++ {
		buf.WriteString(newPathFields[i] + "/")
	}
	buf.WriteString(newPathFields[idx-1])
	return buf.String()
}
