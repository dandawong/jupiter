package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{
	W, H int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0,0, img.W, img.H)
}

func (img Image) At(x, y int) color.Color {
	v := uint8(((x + y) / 2) % 255)
	return color.RGBA{v,v,255,255}
}

func main() {
	m := Image{100,100}
	pic.ShowImage(m)
}
