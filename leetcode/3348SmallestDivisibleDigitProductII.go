package leetcode

import "strings"

func smallestNumber3348(num string, t int64) string {
	temp := t
	for i := int64(2); i <= 9; i++ {
		for temp%i == 0 {
			temp /= i
		}
	}
	if temp > 1 {
		return "-1"
	}

	n := len(num)
	rem := make([]int64, n+1)
	rem[0] = t
	pos := n - 1

	numBytes := []byte(num)
	for i := 0; i < n; i++ {
		if numBytes[i] == '0' {
			pos = i
			break
		}
		rem[i+1] = rem[i] / gcd3348(rem[i], int64(numBytes[i]-'0'))
	}

	if rem[n] == 1 {
		return num
	}

	for i := pos; i >= 0; i-- {
		for {
			numBytes[i]++
			if numBytes[i] > '9' {
				break
			}
			tNow := rem[i] / gcd3348(rem[i], int64(numBytes[i]-'0'))
			k := 9
			for j := n - 1; j > i; j-- {
				for tNow%int64(k) != 0 {
					k--
				}
				tNow /= int64(k)
				numBytes[j] = byte('0' + k)
			}
			if tNow == 1 {
				return string(numBytes)
			}
		}
	}

	var ans strings.Builder
	t = t
	for i := 9; i > 1; i-- {
		for t%int64(i) == 0 {
			ans.WriteByte(byte('0' + i))
			t /= int64(i)
		}
	}

	ansStr := ans.String()
	padding := max(n+1-len(ansStr), 0)
	ansStr += strings.Repeat("1", padding)
	runes := []rune(ansStr)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func gcd3348(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
