package leetcode

import "fmt"

func fizzBuzz(n int) []string {
	r := make([]string, 0)
	for idx := 1; idx <= n; idx++ {
		item := ""
		if idx%3 == 0 {
			item += "Fizz"
		}
		if idx%5 == 0 {
			item += "Buzz"
		}
		if item == "" {
			item = fmt.Sprintf("%d", idx)
		}
		r = append(r, item)
	}
	return r
}
