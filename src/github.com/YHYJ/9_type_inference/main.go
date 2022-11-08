/* File: main.go */
/* Author: YJ */
/* Email: yj1516268@outlook.com */
/* Created Time: 2020-01-16 11:16:53 */

// Description: 类型推导
// 当定义变量不指定类型时（使用`var`不给出类型或者使用`:=`），变量的类型由右值推导得出
// 当右值定义了类型时，新变量的类型与其相同

package main

import (
	"fmt"
)

func main() {
	// 当右值为数字常量且未指定类型时，新的变量的类型取决于常量的精度
	i := 42           // int
	f := 3.0          // float64
	g := 0.867 + 0.5i // complex128
	var x int32 = 12  // int32
	y := x            // int32
	fmt.Printf("Type: %T\n", i)
	fmt.Printf("Type: %T\n", f)
	fmt.Printf("Type: %T\n", g)
	fmt.Printf("Type: %T\n", x)
	fmt.Printf("Type: %T\n", y)
}
