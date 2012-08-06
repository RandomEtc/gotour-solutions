package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	m := make([][]uint8,dy)
	for y := 0; y < dy; y++ {
		m[y] = make([]uint8,dx)
		for x := 0; x < dy; x++ {
			m[y][x] = uint8(x*y)
		}		
	}
	return m
}

func main() {
	pic.Show(Pic)
}