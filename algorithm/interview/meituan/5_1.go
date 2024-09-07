package meituan

import "fmt"

/*
输入一个数组，每次对选择任意一个数字 a[i],执行+x操作(x>0), 花费的代价是 a[i]+x, 求数字变成全是奇数或者全是偶数的最小代价
*/
func problem1() {
	var n, t int
	fmt.Scanf("%d", &n)
	o, e := 0, 0
	so, se := 0, 0
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &t)
		if t&1 == 1 {
			so += t
			o++
		} else {
			e++
			se += t
		}
	}
	if o == 0 || e == 0 {
		fmt.Println(0)
	}
	so += o
	se += e
	if so > se {
		fmt.Println(se)
		return
	}
	fmt.Println(so)
}
