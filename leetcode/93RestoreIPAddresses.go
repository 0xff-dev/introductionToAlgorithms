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
