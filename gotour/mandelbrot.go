package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
	"math/cmplx"
)

type Image struct {
	w, h int
	at   func(x, y, w, h int) int
}

func (this *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (this *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, this.w, this.h)
}

func (this *Image) At(x, y int) color.Color {
	c := uint8(this.at(x, y, this.w, this.h))
	return color.RGBA{c, c, 255, 255}
}

func mandelbrot(col, row, width, height int) int {
	const max = 255

	fc := float64(col)
	fr := float64(row)
	fh := float64(width)
	fw := float64(height)

	c := complex((fc-fw/2)*4/fw, (fr-fh/2)*4/fw)
	z := c

	iter := 0
	for ; cmplx.Abs(z) <= 2 && iter < max; iter++ {
		z = z*z + c
	}

	if iter < max {
		return iter
	} else {
		return 0
	}
}

func main() {
	//sq := func(x, y, w, h int) int { return x * y }
	pic.ShowImage(&Image{500, 500, mandelbrot})
}
