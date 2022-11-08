/* File: main.go */
/* Author: YJ */
/* Email: yj1516268@outlook.com */
/* Created Time: 2020-01-16 10:15:41 */

// Description: 函数
// 函数可以有0~N个参数

package main

import (
	"fmt"
	"math"
)

// function add一次接受两个float64类型的参数
func add(x, y float64) {
	num := x + y
	fmt.Println("x + y =", num)
}

// 函数也可以作为值传递
// compute函数的参数fn是一个未命名函数
// compute函数的返回值是函数fn的执行结果
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	add(3, 3.16)

	// 将一个未命名函数赋值给变量hypot
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Printf("hypot: Type=%T\n", hypot)
	fmt.Printf("hypot: Type=%T\n", compute)
	fmt.Println(hypot(5, 12)) // 以调用函数的形式调用hypot,因为此时hypot的值就是一个函数

	fmt.Println(compute(hypot))    // 将hypot代表的函数作为参数传递给compute函数，因为compute函数返回固定的fn(3, 4)，所以hypot接收3和4，返回5
	fmt.Println(compute(math.Pow)) // func math.Pow(x, y, float64) float64 返回x**y
}
