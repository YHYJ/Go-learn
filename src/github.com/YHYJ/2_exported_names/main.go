/* File: main.go */
/* Author: YJ */
/* Email: yj1516268@outlook.com */
/* Created Time: 2020-01-16 10:14:25 */

// Description: 导出名
// 在导入一个包后，就可以用其导出的名称来调用它
// 在 Go 中，首字母大写的名称是被导出的
// 例如 Pi 就是被导出的，它导出自 math 包
// 而 pi 并未以大写字母开头，所以是不会被导出的
// 任何不被导出的名字在该包外均无法访问

package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(math.Pi)
	fmt.Println(math.Pi)
}
