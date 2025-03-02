package bitmap

import (
	"image"
	"image/color"
)

type Bitmap struct {
	Width, Height int
	Pixels        [][]color.NRGBA // https://stackoverflow.com/a/54309181
}

func pixel2rgba(pixel color.Color) color.NRGBA {
	r, g, b, a := pixel.RGBA()
	if a == 0xffff {
		return color.NRGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), 0xff}
	}
	if a == 0 {
		return color.NRGBA{0, 0, 0, 0}
	}
	// Since Color.RGBA returns an alpha-premultiplied color, we should have r <= a && g <= a && b <= a.
	r = (r * 0xffff) / a
	g = (g * 0xffff) / a
	b = (b * 0xffff) / a
	return color.NRGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
}

func FromImage(img image.Image) Bitmap {
	bounds := img.Bounds()
	height := bounds.Dy()
	width := bounds.Dx()

	bitmap := Bitmap{
		Width:  width,
		Height: height,
		Pixels: make([][]color.NRGBA, height),
	}

	for y := 0; y < height; y++ {
		bitmap.Pixels[y] = make([]color.NRGBA, width)
		for x := 0; x < width; x++ {
			pixel := img.At(x, y)
			bitmap.Pixels[y][x] = pixel2rgba(pixel)
		}
	}

	return bitmap
}
