package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var out [][]uint8
	
	for y := 0; y < dy; y++ {
		out = append(out, make([]uint8, dx))
	}

	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			out[x][y] = uint8(x^y)
		}
	}
	
	return out
}

func main() {
	pic.Show(Pic)
}
