/* File: main.go */
/* Author: YJ */
/* Email: yj1516268@outlook.com */
/* Created Time: 2020-01-13 14:50:23 */

// Description: 包
// 每个go程序都是由包组成的
// 程序运行的入口是包'main'
// 一个代码文件归属于哪个包由第一行非注释的代码`package <name>`决定，归属于同一个包的文件一般放在同一个文件夹下
// 按惯例，包名与导入路径的最后一个目录名一致。
// 例如："fmt"归属于fmt包，"math/rand"归属于rand包

package main

import ( // 导入以下3个包
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		speed int64 = time.Now().UnixNano()
		n     int   = 100
	)
	rand.Seed(speed)                                   // 设置一个随机种子，Init才会每次生成不一样的数字
	fmt.Println("My favorite number is", rand.Intn(n)) // 生成[0, n)的伪随机数
}
