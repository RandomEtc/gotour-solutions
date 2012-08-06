package main

import (
	"image"
	"image/color"
	"code.google.com/p/go-tour/pic"
)

type Image struct{
	w,h int
	px [][]uint8
	model color.Model
}

func (img Image) At(x,y int) color.Color {
	v := img.px[y][x]
	return color.RGBA{v, v, 255, 255}
}

func (img Image) ColorModel() color.Model {
	return img.model
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0,0,img.w,img.h)
}

func main() {
	m := Image{w:1024,h:1024,model:color.RGBAModel}
	m.px = make([][]uint8, m.w)
	for y := 0; y < m.w; y++ {
		m.px[y] = make([]uint8, m.h)
		for x := 0; x < m.h; x++ {
			m.px[y][x] = uint8(x+y)
		}
	}	
	pic.ShowImage(&m)
}