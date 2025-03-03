package bitmap

import (
	"image/color"
)

func pixelToGrayScale(pixel color.NRGBA) color.Gray {
	// https://en.wikipedia.org/wiki/Relative_luminance
	// Luminance Y = 0.2126 R + 0.7152 G + 0.0722 B
	// Approxiamtion for speed:
	// Y = 0.375 R + 0.5 G + 0.125 B
	// Y = (R + R + R + B + G + G + G + G) / 8
	y := (pixel.R + pixel.R + pixel.R + pixel.B + pixel.G + pixel.G + pixel.G + pixel.G) >> 3
	return color.Gray{uint8(y)}
}

func (bitmap Bitmap) ToGrayscale() [][]color.Gray {
	grayPixels := make([][]color.Gray, bitmap.Height)
	for y := 0; y < bitmap.Height; y++ {
		grayPixels[y] = make([]color.Gray, bitmap.Width)
		for x := 0; x < bitmap.Width; x++ {
			grayPixels[y][x] = pixelToGrayScale(bitmap.Pixels[y][x])
		}
	}
	return grayPixels
}
