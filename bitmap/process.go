package bitmap

import (
	"image/color"
)

func pixelToGrayScale(pixel color.NRGBA) uint8 {
	return uint8((uint32(pixel.R) + uint32(pixel.G) + uint32(pixel.B)) / 3)
}

func ToGrayscale(bitmap Bitmap) [][]uint8 {
	grayPixels := make([][]uint8, bitmap.Height)
	for y := 0; y < bitmap.Height; y++ {
		grayPixels[y] = make([]uint8, bitmap.Width)
		for x := 0; x < bitmap.Width; x++ {
			grayPixels[y][x] = pixelToGrayScale(bitmap.Pixels[y][x])
		}
	}
	return grayPixels
}

func ToBrightnessMap(bitmap Bitmap) [][]uint8 {
	brightnessMap := make([][]uint8, bitmap.Height)
	for y := 0; y < bitmap.Height; y++ {
		brightnessMap[y] = make([]uint8, bitmap.Width)
		for x := 0; x < bitmap.Width; x++ {
			P := bitmap.Pixels[y][x]
			// https://en.wikipedia.org/wiki/Relative_luminance
			// Luminance Y = 0.2126 R + 0.7152 G + 0.0722 B
			// Approxiamtion for speed:
			// Y = 0.375 R + 0.5 G + 0.125 B
			// Y = (R + R + R + B + G + G + G + G) / 8
			brightnessMap[y][x] = (P.R + P.R + P.R + P.B + P.G + P.G + P.G + P.G) >> 3
		}
	}
	return brightnessMap
}
