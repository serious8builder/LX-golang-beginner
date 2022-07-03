package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	Width, Height int
	color         uint8
}

func (img *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img *Image) At(x, y int) color.Color {
	return color.RGBA{R: img.color + uint8(x), G: img.color + uint8(y), B: 255, A: 255}
}

func main() {
	m := Image{100, 100, 128}
	pic.ShowImage(&m)
}
