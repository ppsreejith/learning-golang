//gotour
//http://127.0.0.1:3999/moretypes/1

package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
