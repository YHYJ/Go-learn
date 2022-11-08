/* File: main.go */
/* Author: YJ */
/* Email: yj1516268@outlook.com */
/* Created Time: 2020-01-16 11:15:45 */

// Description: 类型转换
// 表达式 `T(v)` 将值 'v' 转换为类型 'T'
// go进行显式的类型转换
// 类型转换最好使用strconv包

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var x, y int = 3, 4
	var f = math.Sqrt(float64(x*x + y*y))
	var z = uint(f)
	fmt.Println(x, y, z, f)
	fmt.Println(string(6))
	fmt.Println(strconv.Itoa(6))
}
