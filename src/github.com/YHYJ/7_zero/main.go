/* File: main.go */
/* Author: YJ */
/* Email: yj1516268@outlook.com */
/* Created Time: 2020-01-16 11:14:59 */

// Description: 零值
// 变量在定义时没有明确的初始化的话会被赋值为对应的'零值'
// 数值类型的零值为0
// 布尔类型的零值为false
// 字符串类型的零值为空字符串
// ......

package main

import (
	"fmt"
)

func main() {
	var (
		i int
		f float64
		b bool
		s string
		bt byte
		r rune
		c complex64
	)
	fmt.Printf("%v %v %v %v %v %v %v\n", i, f, b, s, bt, r, c)
}
