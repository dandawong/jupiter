package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dx)

	for i := range pic {
		pic[i] = make([]uint8, dy)
	}

	for ix := range pic {
		for iy := range pic[ix] {
			pic[ix][iy] = uint8(ix ^ iy)
		}
	}

	return pic
}

func main() {
	pic.Show(Pic)
}
