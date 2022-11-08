/* File: main.go */
/* Author: YJ */
/* Email: yj1516268@outlook.com */
/* Created Time: 2020-01-16 10:16:43 */

// Description: 返回值
// 多返回值：函数可以返回任意数量的返回值，只需在声明函数的时候注意
// 命名返回值：go的返回值可以命名，并且像变量一样使用，返回值名称应具有意义以便作为文档使用

package main

import (
	"fmt"
)

// 一个函数可以有多个返回值，在声明的时候用()括起来
func swap(x, y string) (string, string) {
	return x, y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	a, b := swap("X", "Y")
	fmt.Printf("a = %s, b = %s\n", a, b)
	fmt.Println(split(19))
}
