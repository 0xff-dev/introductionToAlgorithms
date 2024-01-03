package leetcode

import (
	"regexp"
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	if isIPv4(IP) {
		return "IPv4"
	} else if isIPv6(IP) {
		return "IPv6"
	}
	return "Neither"
}

func isIPv4(IP string) bool {
	pattern := `^(([1-9]\d{0,2}|0)\.){3}([1-9]\d{0,2}|0)$`
	match, _ := regexp.MatchString(pattern, IP)
	if !match {
		return false
	}

	segments := strings.Split(IP, ".")
	for _, segment := range segments {
		num, _ := strconv.Atoi(segment)
		if num > 255 || (segment[0] == '0' && len(segment) > 1) {
			return false
		}
	}

	return true
}

func isIPv6(IP string) bool {
	pattern := `^([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}$`
	match, _ := regexp.MatchString(pattern, IP)
	return match
}
