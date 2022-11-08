/* File: main.go */
/* Author: YJ */
/* Email: yj1516268@outlook.com */
/* Created Time: 2020-01-16 15:48:04 */

// Description: 闭包
// go函数可以闭包
// 闭包是一个作为返回值的函数，它引用了外部变量

package main

import (
	"fmt"
)

func adder() func(int) int {
	sum := 0                 // 闭包外部变量sum初始化等于0
	return func(x int) int { // adder函数返回一个闭包
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	sum := 1 // 闭包外部变量sum初始化等于0
	sun := 0
	return func() int { // adder函数返回一个闭包
		sum, sun = sun, sum+sun
		return sum
	}
}

func main() {
	pos, neg := adder(), adder() // 两个外部变量sum是独立的
	fmt.Printf("pos Type: %T\n", pos)
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-i))
	}

	sun := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(sun())

	}
}
