package bitmap

import (
	"image/color"
	"math"
)

func (b *Bitmap) Resize(width int) *Bitmap {
	//Width should be greater than 0 and less than the original width
	if width <= 0 || width >= b.Width {
		return b
	}

	aspectRatio := float64(b.Height) / float64(b.Width)
	height := int(float64(width) * aspectRatio)

	resized := Bitmap{
		Width:  width,
		Height: height,
		Pixels: make([][]color.NRGBA, height),
	}

	for y := 0; y < height; y++ {
		resized.Pixels[y] = make([]color.NRGBA, width)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			srcX := float64(x) * float64(b.Width) / float64(width)
			srcY := float64(y) * float64(b.Height) / float64(height)

			x0, y0 := int(srcX), int(srcY)
			x1, y1 := int(math.Min(float64(x0+1), float64(b.Width-1))), int(math.Min(float64(y0+1), float64(b.Height-1)))

			weightX := srcX - float64(x0)
			weightY := srcY - float64(y0)

			c00 := b.Pixels[y0][x0]
			c01 := b.Pixels[y0][x1]
			c10 := b.Pixels[y1][x0]
			c11 := b.Pixels[y1][x1]

			r, g, b, a := interpolateColors(c00, c01, c10, c11, weightX, weightY)

			resized.Pixels[y][x] = color.NRGBA{R: r, G: g, B: b, A: a}
		}
	}
	return &resized
}

func interpolateColors(c00, c01, c10, c11 color.NRGBA, weightX, weightY float64) (uint8, uint8, uint8, uint8) {
	r := interpolate(c00.R, c01.R, c10.R, c11.R, weightX, weightY)
	g := interpolate(c00.G, c01.G, c10.G, c11.G, weightX, weightY)
	b := interpolate(c00.B, c01.B, c10.B, c11.B, weightX, weightY)
	a := interpolate(c00.A, c01.A, c10.A, c11.A, weightX, weightY)

	return r, g, b, a
}

func interpolate(c00, c01, c10, c11 uint8, weightX, weightY float64) uint8 {
	// Horizontal interpolation
	c0 := float64(c00)*(1-weightX) + float64(c01)*weightX
	c1 := float64(c10)*(1-weightX) + float64(c11)*weightX

	// Vertical interpolation
	c := c0*(1-weightY) + c1*weightY

	return uint8(math.Round(c))
}
