/* File: main.go */
/* Author: YJ */
/* Email: yj1516268@outlook.com */
/* Created Time: 2020-01-16 15:18:35 */

// Description: 数据类型`array`
// 数组的长度是其类型的一部分，因此数组不能改变大小
// 伪代码 `array1 := [length]type` 表示：
// array1是一个数组，该数组有length个元素，元素类型为type

package main

import (
	"fmt"
)

func main() {
	var a [2]string // 定义数组a，有2个string类型的元素

	// 使用下标定义数组元素
	a[0] = "Hello"
	a[1] = "World"
	// 使用下标访问数组元素
	fmt.Println(a)
	fmt.Println(a[0], a[1])

	primes := [6]int{1, 2, 3, 4, 5, 6} // 定义数组primes，有6个int类型的元素

	fmt.Println(primes)
}
