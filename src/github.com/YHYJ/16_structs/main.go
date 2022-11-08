/* File: main.go */
/* Author: YJ */
/* Email: yj1516268@outlook.com */
/* Created Time: 2020-01-16 15:17:15 */

// Description: 数据类型`struct`
// 一个结构体是一组字段(field)的集合

package main

import (
	"fmt"
)

/* 结构体的定义 */
// Vertex 是新定义的一个结构体
type Vertex struct {
	// 结构体 Vertex 的5个元素
	A, B int
	X    uint
	Y    float32
	Z    string
}

/* 结构体的初始化 */
var (
	// 1. 按元素名赋值，返回对象
	// 未赋值的元素被隐式赋予零值
	v = Vertex{
		A: 1, B: 2,
		X: 190, Y: 22,
		Z: "i am struct",
	}
	v1 = Vertex{X: 1, A: 12} // 字段名顺序无关
	v2 = Vertex{}

	// 2. 按位置赋值，返回对象
	// 必须全部显式赋值
	v3 = Vertex{10, 20, 1900, 220, "location"}

	// 3. 返回指针
	v4 = new(Vertex)
)

func main() {
	fmt.Println(v)
	// 使用'struct.field'访问结构体元素
	fmt.Println(v.X)
	fmt.Println(v.Y)
	fmt.Println(v.Z)

	fmt.Println("----------------------------------")

	// 使用指针访问结构体
	V := &v
	// 注意结构体指针不同于普通指针
	// 为了允许隐式间接引用，这里V的输出内容做了调整
	// 没有直接输出内存地址，而是输出了取地址符&, 但实质上输出的还是内存地址
	fmt.Printf("V -- Type: %T, Value: %v\n", V, V)
	V.X = 110
	fmt.Println((*V).X, V.X) // 和普通指针一样的获取内存空间里的数据的方法，但因为这样写太啰嗦，所以go允许隐式间接引用
	fmt.Printf("V -- Type: %T, Value: %v\n", *V, *V)

	fmt.Println("----------------------------------")

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(v4)
}
