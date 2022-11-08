/* File: main.go */
/* Author: YJ */
/* Email: yj1516268@outlook.com */
/* Created Time: 2020-01-16 15:19:16 */

// Description: 数据类型`slice`
// 每个`slice`都有一个底层数组
// `slice`不存储任何数据，它只是描述了其底层数组的一部分
// 更改`slice`的元素对于其底层数组和其他引用是同步的
// `slice`拥有“长度len”和“容量cap”：`slice`的长度就是其元素数量；容量就是其底层数组的元素数量
// 伪代码 slice1 := []type 表示：
// slice1是一个slice，其元素类型为type
// 切片操作和数据类型`slice`不一样

package main

import (
	"fmt"
	"strings"
)

func slices() {
	var slice1 = []int{1, 2, 3, 4, 5}
	// 定义`slice`：该行会创建一个`array`=[5]int{1, 2, 3, 4, 5}，然后构建一个引用了它的切片
	slice1 = slice1[:0] // 因为`slice`都是对底层数组的引用
	slice1 = slice1[:4] // 所以每次对`slice`的切片都是对底层数组的切片
	slice1 = slice1[2:] // 所以这3个切片操作并不会互相影响
	fmt.Printf("Type(slice1) = %T\n", slice1)
	fmt.Printf("Value(slice1) = %v\n", slice1)
	fmt.Printf("len(slice1) = %v\n", len(slice1))
	structSlice := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{10, true},
		{7, true},
	}
	fmt.Printf("Type(structSlice) = %T\n", structSlice)
	fmt.Printf("Value(structSlice) = %v\n", structSlice)
	fmt.Printf("len(structSlice) = %v\n", len(structSlice))
}

// 可以对`slice`切片
func slicingSlices() {
	slice2 := []int{1, 3, 5, 7, 9}
	fmt.Println("slices2 =", slice2)
	fmt.Println("slices2[1:3] =", slice2[1:3])
	fmt.Println("slices2[:3] =", slice2[:3]) // 省略第一个下标代表从下标0开始切片
	fmt.Println("slices2[1:] =", slice2[1:]) // 省略第二个下标代表从切片到最后一个下标结束
	// go的切片操作不支持负数和倒序切片
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

// `slice`由函数`make`创建，这会分配一个指定长度的数组并且返回一个`slice`指向这个数组
// `make`的第3个参数指定容量(cap)，因此限制了`slice`支持的最大切片长度
func makingSlices() {
	a := make([]int, 5) // 创建一个`slice`，len = 5，cap默认等于len
	a[4] = 1            // len = 5，所以最大下标为4
	fmt.Printf("a = %v -- len=%d, cap=%d\n", a, len(a), cap(a))
	b := make([]int, 1, 6) // 创建一个`slice`，长度为1
	b[0] = 123             // len = 1，所以最大下标为0
	fmt.Printf("b = %v -- len=%d, cap=%d\n", b, len(b), cap(b))
	c := b[:6] // len <= cap
	fmt.Printf("c = %v -- len=%d, cap=%d\n", c, len(c), cap(c))
}

// `slice`的零值是`nil` —— 这里指的是`slice`本身的零值，而非`slice`的元素的零值
// `nil slice`的`len`和`cap`都是0且没有底层数组
func nilSlice() {
	var z []int
	fmt.Printf("type=%T, len=%d, cap=%d\n", z, len(z), cap(z))
	if z == nil {
		fmt.Println("z is nil")
	}
}

// 向`slice`添加元素
func appendSlice() {
	var a []int
	fmt.Println(a)

	a = append(a, 0) // append 返回更新后的slice
	fmt.Println(a)

	a = append(a, 1, 2, 3)
	fmt.Println(a)
}

func main() {
	slices()
	fmt.Println("-------------------------------")
	slicingSlices()
	fmt.Println("-------------------------------")
	makingSlices()
	fmt.Println("-------------------------------")
	nilSlice()
	fmt.Println("-------------------------------")
	appendSlice()
}
