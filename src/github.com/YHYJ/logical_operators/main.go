/* File: main.go */
/* Author: YJ */
/* Email: yj1516268@outlook.com */
/* Created Time: 2020-01-16 15:40:07 */

// Description: 逻辑运算符

package main

import (
	"fmt"
)

func main() {
	var x bool = true
	var y bool = false

	fmt.Println("x =", x)
	fmt.Println("y =", y)

	fmt.Println("x && y =", x && y) // and
	fmt.Println("x || y =", x || y) // or
	fmt.Println("!x =", !x)         // not
}
