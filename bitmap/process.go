package bitmap

import (
	"image/color"
)

func pixelToGrayScale(pixel color.NRGBA, invert bool) color.Gray {
	gray := color.GrayModel.Convert(pixel).(color.Gray)
	if invert {
		gray.Y = 255 - gray.Y
	}
	return gray
}

func (bitmap Bitmap) ToGrayscale(invert bool) [][]color.Gray {
	grayPixels := make([][]color.Gray, bitmap.Height)
	for y := 0; y < bitmap.Height; y++ {
		grayPixels[y] = make([]color.Gray, bitmap.Width)
		for x := 0; x < bitmap.Width; x++ {
			grayPixels[y][x] = pixelToGrayScale(bitmap.Pixels[y][x], invert)
		}
	}
	return grayPixels
}