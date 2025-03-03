package bitmap

import (
	"image/color"
)

func pixelToGrayScale(pixel color.NRGBA) color.Gray {
	return color.GrayModel.Convert(pixel).(color.Gray)
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
