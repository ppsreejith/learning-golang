//gotour
//http://127.0.0.1:3999/methods/1

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prev, curr := 0, 1
	return func() int {
		defer func() {
			prev, curr = curr, curr+prev
		}()
		return prev
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
