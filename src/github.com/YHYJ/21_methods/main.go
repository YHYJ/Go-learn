/* File: main.go */
/* Author: YJ */
/* Email: yj1516268@outlook.com */
/* Created Time: 2020-01-16 15:24:04 */

// Description: method
// go没有类(class)的概念，但可以定义'方法(method)'
// 方法是具有特殊的'方法接收器(receiver)'参数的函数（本质上还是一个函数）
// '方法接收器'作为独立参数列表出现在关键字`func`和方法(method)名之间

package main

import (
	"fmt"
	"math"
)

// Vertex 结构体
type Vertex struct {
	X, Y float64 // 'X'和'Y'作为field绑定到结构体Vertex上
}

// VAbs 方法
// 它的接收器是`v`，类型是`Vertex`
func (v Vertex) VAbs() float64 { // 'VAbs'作为method绑定到结构体Vertex上
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Abs 普通函数
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Mfloat 类型
// 自定义类型，本质是一个float64
type Mfloat float64

// MAbs 方法
// 为Mfloat类型定义MAbs方法
func (i Mfloat) MAbs() float64 {
	if i < 0 {
		return float64(-i)
	}
	return float64(i)
}

type mint int64

func (m mint) MInt() int64 {
	return int64(m)
}

func main() {
	v := Vertex{X: 12, Y: 16}
	fmt.Println(v.VAbs())
	fmt.Println(Abs(v))

	m := Mfloat(math.Sqrt2)
	fmt.Println(m.MAbs())
	fmt.Println(math.Sqrt2)
}
