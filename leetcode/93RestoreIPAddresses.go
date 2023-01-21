package leetcode

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	r := make([]string, 0)
	if len(s) < 4 {
		return r
	}
	ipAux(s, 0, 0, "", &r)
	return r
}

func ipAux(s string, idx, dot int, ip string, r *[]string) {
	if dot == 3 {
		if idx >= len(s) {
			return
		}
		if less255(s[idx:]) {
			*r = append(*r, ip+s[idx:])
		}
		return
	}
	if idx >= len(s) {
		return
	}
	if s[idx] == '0' {
		ipAux(s, idx+1, dot+1, ip+"0.", r)
		return
	}
	for i := idx + 1; i <= idx+3 && i < len(s); i++ {
		item := s[idx:i]
		if less255(item) {
			ipAux(s, i, dot+1, ip+item+".", r)
			continue
		}
		return
	}
}

func less255(s string) bool {
	if s == "0" {
		return true
	}
	if strings.HasPrefix(s, "0") {
		return false
	}
	r, _ := strconv.Atoi(s)
	return r <= 255
}

func restoreIpAddresses2(s string) []string {
	ans := make([]string, 0)
	var dfs func(int, int, string)
	dfs = func(start, dot int, prefix string) {
		if start >= len(s) {
			return
		}
		if dot == 0 {
			diff := len(s) - start
			if diff != 1 && s[start] == '0' {
				return
			}
			now := s[start:]

			now1, _ := strconv.Atoi(now)
			if now1 >= 0 && now1 <= 255 {
				ans = append(ans, prefix+now)
			}
			return
		}
		if s[start] == '0' {
			dfs(start+1, dot-1, prefix+"0.")
			return
		}

		for end := start; end <= start+2 && end < len(s); end++ {
			now := s[start : end+1]
			now1, _ := strconv.Atoi(s[start : end+1])
			if now1 >= 0 && now1 <= 255 {
				dfs(end+1, dot-1, prefix+fmt.Sprintf("%s.", now))
			}
		}
	}

	dfs(0, 3, "")
	return ans
}
