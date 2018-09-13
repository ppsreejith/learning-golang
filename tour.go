// gotour
// http://127.0.0.1:3999/basics/7

package main

import (
	"fmt"
)

func main() {
	var (
		ToBe   bool   = false
		maxInt uint64 = 1<<64 - 1
	)
	var (
		i int     = 42
		f float64 = float64(i)
		u uint    = uint(i)
	)
	const World = "Bhoomi"
	fmt.Println(ToBe, maxInt, i, f, u, World)
}
