// gotour
// http://127.0.0.1:3999/basics/7

package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(swap("hello", "world"))
}
