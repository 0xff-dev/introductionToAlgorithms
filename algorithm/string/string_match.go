package string

import "fmt"

func NativeMatch(s, t string) []int{
	indexs := make([]int, 0)
	sl, tl := len(s), len(t)
	if sl > tl {
		return indexs
	}
	for start := 0; start < tl - sl + 1; start ++ {
		i := 0
		for ; i < sl ; i ++ {
			if t[start] != s[i] {
				break
			}
		}
		if i == sl {
			indexs = append(indexs, start)
		}
	}
	return indexs
}

// Rabin-Karp 算法
// 先利用O(m)的时间计算出s 的值，在利用O(n-m+1)的时间计算每个Ts T[s+1:s+m+1]范围的值
// 最后利用O(m)+O(n-m+1) 时间计算出结果
// 取余数，会出现伪命中点，当匹配的时候需要做一下匹配是否是伪命中点
// 目前的实现形式是针对数字字符串，所以如果涉及到字符相关， 需要做下变动。

// return int(s) and base^m-1, 方便后续计算
func beforeRabinKarp(s, t string, base, mod int) (s0, t0, p int) {
	s0, t0, p = 0, 0, 1
	for index := 0; index < len(s); index ++ {
		s0 = (base*s0 + int(byte(s[index])-'0'))%mod
		t0 = (base*t0 + int(byte(t[index]-'p')))%mod
		if index != 0 {
			p *= base
		}
	}
	return
}

func RabinKarp(s, t string) {
	s0, t0, power := beforeRabinKarp(s, t, 10, 13)
	for index := 0; index < len(t)-len(s); index ++ {
		fmt.Println("t0 := ", t0)
		if t0 == s0 {
			// compare
			_inner := 0
			for ; _inner < len(s); _inner ++ {
				if t[index+_inner] != s[_inner] {
					break
				}
			}
			if _inner == len(s) {
				fmt.Println("found start: ", index, " : ", t[index:index+len(s)])
			} else {
				fmt.Println("same mod not match", index, " : ", t[index:index+len(s)])
			}
		} else {
			// 重新计算t0
			t0 = (10*(t0-power*int(t[index]-'0')) + int(t[index+len(s)]-'0')) % 13
		}
	}
}

func nextArray(s string) []int {
	nextArr := make([]int, len(s))
	nextArr[0] = -1
	compareIndex := -1
	index := 0
	for index < len(s)-1 {
		if compareIndex == -1 || s[index] == s[compareIndex] {
			compareIndex++
			index++
			nextArr[index] = compareIndex
			continue
		}
		compareIndex = nextArr[compareIndex]
	}
	return nextArr
}

func kmp(s, t string) {
	next := nextArray(s)
	ls, lt := len(s), len(t)
	fmt.Println("next: ", next)
	i, j := 0, 0
	for i < ls && j < lt {
		if i == -1 || s[i] == t[j] {
			i++
			j++
			continue
		}
		i = next[i]
	}
	if i >= ls {
		fmt.Println("start: ", j-i)
	}
}