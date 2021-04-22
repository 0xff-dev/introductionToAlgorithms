package leetcode

import "fmt"

//func gen() chan int {
//	ch := make(chan int)
//	go func() {
//		for i := 2; ; i++ {
//			ch <- i
//		}
//	}()
//	return ch
//}
//
//
//func filter(in chan int, key int) chan int {
//	r := make(chan int)
//	go func() {
//		for i := range in {
//			if i % key != 0 {
//				r <-i
//			}
//		}
//		close(r)
//	}()
//	return r
//}
func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	//numCh := gen()
	//var (
	//	i int
	//	ok bool
	//)
	//r := 0
	//for i < n {
	//	i, ok = <- numCh
	//	if !ok {
	//		break
	//	}
	//	fmt.Println("now: ", i)
	//	r++
	//	numCh = filter(numCh, i)
	//}
	r := 0
	arr := make([]int, n)
	for i := 2; i < n; i++ {
		if arr[i] == 0 {
			fmt.Println("now: ", i)
			r++
			for j := i + i; j < n; j += i {
				arr[j] = 1
			}
		}
	}
	return r
}
