/* File: main.go */
/* Author: YJ */
/* Email: yj1516268@outlook.com */
/* Created Time: 2020-01-16 10:51:29 */

// Description: go的基本类型：
/*
bool
string
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
byte              // uint8的别名
rune              // int32的别名，代表一个Unicode码
float32 float64   // 小数点后6位；小数点后15位
complex64 complex128
*/
// int, uint 和 uintptr 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽

package main

import (
	"fmt"
	"math/cmplx"
)

var name string = "yj"

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", name, name)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n\n", z, z)
}
