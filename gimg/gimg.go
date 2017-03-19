package gimg

import (
	"fmt"
	"github.com/Lewuathe/gimg/ansi"
	"github.com/nfnt/resize"
	"image"
)

// Resize image into the size which can be printable in terminal.
func ResizeImage(xsize int, src image.Image) image.Image {
	xmax := src.Bounds().Max.X

	ratio := float64(xsize) / float64(xmax)
	width := uint(float64(xmax) * ratio)
	return resize.Resize(width, 0, src, resize.Bicubic)
}

// Print the image by using ANSI x-term color in terminal
func PrintImage(img image.Image) {
	for i := 0; i < img.Bounds().Max.Y; i++ {
		for j := 0; j < img.Bounds().Max.X; j++ {
			r, g, b, _ := img.At(j, i).RGBA()
			r = ansi.To256(r)
			g = ansi.To256(g)
			b = ansi.To256(b)
			ansi.Print(r, g, b)
		}
		fmt.Printf("\n")
	}
}
