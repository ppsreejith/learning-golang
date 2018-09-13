package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (root float64) {
	for root = 1; math.Abs(root*root-x) > 0.001; {
		root -= (root*root - x) / (2 * root)
	}
	return
}

func main() {
	fmt.Println(Sqrt(2))
}
