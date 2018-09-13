//gotour
//http://127.0.0.1:3999/moretypes/1

package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
