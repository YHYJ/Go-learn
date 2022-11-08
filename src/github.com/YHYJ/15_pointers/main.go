/* File: main.go */
/* Author: YJ */
/* Email: yj1516268@outlook.com */
/* Created Time: 2020-01-16 15:16:43 */

// Description: 指针保存一个变量的内存地址
// go没有指针运算
// *T 是指向类型为 T 的值的指针，零值是nil
// & 是取地址符，会生成一个指向其作用对象的内存地址的指针
// * 是取值符，会获取指针指向的内存地址的数据

package main

import (
	"fmt"
)

func main() {
	cache := 5.0
	var x *float64 = &cache // 指针变量x指向的是cache变量的内存地址
	fmt.Printf("Type: %T, Value: %v\n", x, x)
	*x = 6.0 // 通过指针x设置cache
	fmt.Printf("Type: %T, Value: %v\n", cache, cache)
	fmt.Printf("Type: %T, Value: %v\n", x, x)
	fmt.Printf("Type: %T, Value: %v\n", *x, *x) // 通过指针x获取cache的值，既所谓的‘间接引用/非直接引用’
}
