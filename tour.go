//gotour
//http://127.0.0.1:3999/moretypes/13

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) (pic [][]uint8) {
	for i := 0; i < dy; i++ {
		pic = append(pic, make([]uint8, dx, dx))
		for j := 0; j < dx; j++ {
			pic[i][j] = uint8(i * j)
		}
	}
	return
}

func main() {
	pic.Show(Pic)
}
