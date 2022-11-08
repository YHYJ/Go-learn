/* File: main.go */
/* Author: YJ */
/* Email: yj1516268@outlook.com */
/* Created Time: 2020-01-16 15:21:20 */

// Description: range
// `slice`可以对`slice`或`map`进行迭代

package main

import (
	"fmt"
)

func main() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for k, v := range pow {
		fmt.Printf("2**%d = %d\n", k, v)
	}

	pov := make([]int, 10)
	for i := range pov {
		pov[i] = 1 << uint(i)
	}
	fmt.Println(pov)
	for k := range pov { // 只取`slice`的key
		fmt.Printf("key = %d\n", k)
	}
	for _, v := range pov { // 只取`slice`的value
		fmt.Printf("value = %d\n", v)
	}
}
