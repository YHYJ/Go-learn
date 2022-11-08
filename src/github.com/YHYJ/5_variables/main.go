/* File: main.go */
/* Author: YJ */
/* Email: yj1516268@outlook.com */
/* Created Time: 2020-01-16 10:20:02 */

// Description: 变量
// `var` 语句定义变量
// 在函数中，简洁赋值语句`:=`在类型明确时可以提到`var`进行变量声明

package main

import (
	"fmt"
)

// A 变量的声明，不初始化
var A string

// 声明变量并初始化
var x = 1
var q, t = 1, 2

var (
	a = 1
	b = 2
	c = 3.0
)

/* 短声明变量 */
// 函数外的每个语句必须以关键字开始(var、func等)，所以`:=`不能用在函数之外
func main() {
	p := 1 // 短声明变量必须给一个明确类型的值
	fmt.Println(p)
}
