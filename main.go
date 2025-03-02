package main

import (
	"asciigen/io"
	"fmt"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
)

type Bitmap struct {
	Width, Height int
	Pixels        [][]color.RGBA
}

func main() {
	var1 := 1
	var var2 int = 2
	fmt.Println("Hello World", var1, var2)
	filePath := "test.jpg"
	img, _ := io.ReadImageFromPath(filePath)

	bounds := img.Bounds()
	fmt.Println("Width:", bounds.Dx())
	fmt.Println("Height:", bounds.Dy())

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

	fmt.Println("Successfully read image", filePath)
	fmt.Println("First pixel:", bitmap.Pixels[0][0])
}
