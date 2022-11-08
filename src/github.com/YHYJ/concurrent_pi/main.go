/* File: main.go */
/* Author: YJ */
/* Email: yj1516268@outlook.com */
/* Created Time: 2020-01-16 15:35:17 */

// Description: 并发pi计算

// Cocurrent computation of pi.
// See https://goo.gl/la6Kli
//
// This demonstrates Go's ability to handle large numbers of concurrent processes
// It is an unreasonable way to calculate pi.

package main

import (
	"fmt"
	"math"
)

func term(ch chan float64, k float64) {
	ch <- 4 * math.Pow(-1, k) / (2*k + 1)
}

// pi launches n goroutines to compute an approximation of pi.
func pi(n int) float64 {
	ch := make(chan float64)
	for k := 0; k <= n; k++ {
		go term(ch, float64(k))
	}
	f := 0.0
	for k := 0; k <= n; k++ {
		f += <-ch
	}
	return f
}

func main() {
	fmt.Println(pi(5000))
}
