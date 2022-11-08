/* File: main.go */
/* Author: YJ */
/* Email: yj1516268@outlook.com */
/* Created Time: 2020-01-16 15:14:10 */

// Description: `switch`条件语句
// go的`switch`语句只运行第一个匹配的`case`，因为go自动在每个`case`后面隐式添加了`break`语句
// 除非以`fallthrough`语句结束，否则分支会自动终止
// go的`switch`的`case`无需为常量，且取值不必为整数
// 有条件的`switch`根据条件和`case`是否匹配执行相应分支
// 没有条件的`switch`根据`case`是否成立执行相应分支

package main

import (
	"fmt"
	"runtime"
	"time"
)

func baseSwitch() {
	// 只要匹配到一个case就不再执行后面的case包括default
	// default用来应对所有case都不匹配的情况
	fmt.Println("Go runs on: ")
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Current system is Linux.")
	case "darwin":
		fmt.Println("Current system is macOS.")
	default:
		fmt.Printf("Current system is %v?\n", os)
	}
}

func switchWithNoCondition() {
	// 没有条件的`switch` == `switch true`
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}

func typeSwitch() {
	// 根据变量的数据类型执行相应的动作
	var x interface{}
	x = 1.2

	switch t := x.(type) {
	case nil:
		fmt.Printf("x 的数据类型: %T\n", t)
	case int:
		fmt.Printf("x 的数据类型: %T\n", t)
	case string:
		fmt.Printf("x 的数据类型: %T\n", t)
	case float32, float64:
		fmt.Printf("x 的数据类型: %T\n", t)
	}
}

func main() {
	baseSwitch()
	switchWithNoCondition()
	typeSwitch()
}
