package leetcode

func removeComments(source []string) []string {
	if len(source) == 0 {
		return source
	}

	r := make([]string, 0)
	var tmpBytes []byte
	codeSnippets := false
	for _, line := range source {
		if !codeSnippets {
			tmpBytes = []byte{}
		}
		idx := 0
		for idx < len(line) {
			next := idx + 2
			if next >= len(line) {
				next = len(line)
			}
			if line[idx:next] == "/*" && !codeSnippets {
				codeSnippets = true
				idx++
			} else if line[idx:next] == "*/" && codeSnippets {
				codeSnippets = false
				idx++
			} else if !codeSnippets && line[idx:next] == "//" {
				break
			} else if !codeSnippets {
				tmpBytes = append(tmpBytes, line[idx])
			}
			idx++
		}
		if len(tmpBytes) > 0 && !codeSnippets {
			r = append(r, string(tmpBytes))
		}
	}
	return r
}
