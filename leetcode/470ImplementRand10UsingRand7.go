package leetcode

import "math/rand"

func rand7() int {
	return rand.Intn(7) + 1
}
func rand10() int {
	for {
		a := rand7()
		b := rand7()
		num := (a-1)*7 + b
		if num <= 40 {
			return (num-1)%10 + 1
		}
	}
}
