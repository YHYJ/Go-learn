/* File: main.go */
/* Author: YJ */
/* Email: yj1516268@outlook.com */
/* Created Time: 2020-01-16 15:22:57 */

// Description: map
// `map`将 'key -> value' 进行映射
// `map`的零值是`nil`，`nil map`没有key，也不可以添加key
// `make`函数返回给定类型的`map`，该`map`已初始化
// 伪代码 make(map(type1)type2) 表示：
// 初始化一个`map`，该`map`的key类型是type1，value类型是type2

package main

import (
	"fmt"
)

func basicMap() {
	type Vertex struct {
		Lat, Long float64
	}

	var m = make(map[string]Vertex)
	var n = make(map[string]int)
	var y = map[string]Vertex{}
	var x = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	m["Bell Labs"] = Vertex{Lat: 40.68433, Long: -74.39967}
	fmt.Printf("m = %v -- Type: %T\n", m, m)
	fmt.Println(m["Bell Labs"])

	n["Master"] = 12
	n["Slave"] = -12
	n["pb"] = 1e1
	fmt.Printf("n = %v -- Type: %T\n", n, n)
	fmt.Println(n["Master"])
	fmt.Println(n["Slave"])
	fmt.Println(x)
	fmt.Println(y)
}

func mutatingMap() {
	m := make(map[string]int)
	fmt.Println(m)

	m["Answer"] = 42 // 向`map`插入元素
	m["Question"] = 42
	fmt.Println(m)

	m["Answer"] = 11 // 更新`map`的元素
	fmt.Println(m)

	value := m["Answer"] // 获取`map`元素的值
	fmt.Println(value)

	delete(m, "Answer") // 删除`map`的元素
	fmt.Println(m)

	v, ok := m["Answer"] // 测试指定的key是否存在
	fmt.Println(v, ok)   // 不存在，v=0（零值）；ok=false
	v, ok = m["Question"]
	fmt.Println(v, ok) // 存在，v=value（相应值）；ok=true
}

func main() {
	basicMap()
	fmt.Println("------------------------------------------")
	mutatingMap()
}
