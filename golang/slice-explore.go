package golang

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
   golang的slice作为函数参数传值，但是操作的底层是同一个数组
   24-27 访问超过的元素

   33-36
   不推荐使用 d := append(c, x), 可能产生歧义
*/

func run() {
	a := []int{1, 2}

	fmt.Println("len: ", len(a), " cap:", cap(a))
	a = append(a, 3)
	fmt.Println("len :", len(a), " cap: ", cap(a))
	append1(a)
	rfa := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	fmt.Printf("main slice hader %+v\n", rfa)

	ip := &a[2]
	address := uintptr(unsafe.Pointer(ip)) + unsafe.Sizeof(a[0])
	ip = (*int)(unsafe.Pointer(address))
	fmt.Println("main unsafe address: ", *ip)

	c := []int{1, 2}
	d := append(c, 3)
	c = append(c, 4)
	fmt.Println("c: ", c, " ----> d: ", d)
}

func append1(a []int) {
	fmt.Println("here is appedn1")
	fmt.Println("len :", len(a), " cap: ", cap(a))
	//a[0] = 18
	fmt.Printf("%#v\n appedn 9999\n", a)
	b := append(a, 999999)
	rfa := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	fmt.Printf("slice hader rfa: %+v\n", rfa)

	ip := &a[2]
	address := uintptr(unsafe.Pointer(ip)) + unsafe.Sizeof(a[0])
	ip = (*int)(unsafe.Pointer(address))
	fmt.Println("unsafe address: ", *ip)

	rfb := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	fmt.Printf("slice header rfb: %+v\n\n", rfb)
}
