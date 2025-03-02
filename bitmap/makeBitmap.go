package bitmap

import (
	"image"
	"image/color"
)

type Bitmap struct {
	Width, Height int
	Pixels        [][]color.RGBA
}

func FromImage(img image.Image) Bitmap {
	bounds := img.Bounds()
	height := bounds.Dy()
	width := bounds.Dx()

	bitmap := Bitmap{
		Width:  width,
		Height: height,
		Pixels: make([][]color.RGBA, height),
	}

	for y := 0; y < height; y++ {
		bitmap.Pixels[y] = make([]color.RGBA, width)
		for x := 0; x < width; x++ {
			pixel := img.At(x, y)
			r, g, b, a := pixel.RGBA()
			bitmap.Pixels[y][x] = color.RGBA{uint8(r / 256), uint8(g / 256), uint8(b / 256), uint8(a / 256)}
		}
	}

	return bitmap
}
